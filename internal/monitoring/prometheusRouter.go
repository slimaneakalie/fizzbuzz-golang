package monitoring

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

const (
	_httpRequestTotalMetricName     = "http_request_total"
	_pathLabelName                  = "path"
	_jsonQueryLabelName             = "json_query"
	_responseStatusLabelName        = "response_status"
	_prometheusQueryEndpoint        = "/api/v1/query"
	_prometheusVectorQueryParamName = "query"
)

var (
	_prometheusResponseParseError = errors.New("GetMostFrequentQuery: can't parse response from prometheus")
)

func NewPrometheusHandler(prometheusHostName string) Handler {
	return &PrometheusHandler{
		prometheusVectorQueryURL: fmt.Sprintf("%s%s", prometheusHostName, _prometheusQueryEndpoint),
		promHttpHandler:          promhttp.Handler(),
		httpRequestsCounterVector: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: _httpRequestTotalMetricName,
			Help: "The total number of http requests",
		}, []string{_pathLabelName, _jsonQueryLabelName, _responseStatusLabelName}),
	}
}

func (handler *PrometheusHandler) HandleMonitoringQuery() fizzhttp.HandlerFunc {
	return handler.monitoringQueryHandler
}

func (handler *PrometheusHandler) monitoringQueryHandler(context fizzhttp.RequestContext) {
	handler.promHttpHandler.ServeHTTP(context.GetResponseWriter(), context.GetRequest())
}

func (handler *PrometheusHandler) MonitoringMiddleWare() fizzhttp.HandlerFunc {
	return func(context fizzhttp.RequestContext) {
		context.Next()
		path := context.GetRequest().URL.Path
		jsonStrQuery := context.GetJsonStringQuery()
		responseStatus := strconv.Itoa(context.GetResponseStatus())

		handler.httpRequestsCounterVector.WithLabelValues(path, jsonStrQuery, responseStatus).Inc()
	}
}

func (handler *PrometheusHandler) GetMostFrequentQuery(path string, responseStatus int) (queryRawData *MostFrequentQueryRawData, err error) {
	prometheusQuery := buildMostFrequentRequestPrometheusQuery(path, responseStatus)
	vectorResult, prometheusErr := performPrometheusHttpVectorQuery(prometheusQuery, handler.prometheusVectorQueryURL)
	if prometheusErr != nil {
		return nil, prometheusErr
	}

	queryRawData = &MostFrequentQueryRawData{
		RawStrQuery:  "{}",
		NumberOfHits: 0,
	}

	conversionErr := convertVectorResultToMostFrequentQueryRawData(vectorResult, queryRawData)
	if conversionErr != nil {
		return nil, conversionErr
	}

	return queryRawData, nil
}

func buildMostFrequentRequestPrometheusQuery(path string, responseStatus int) string {
	// it creates something like topk(1, sum by (json_query) (http_request_total{path="/v1/fizzbuzz", response_status="200"}))
	return fmt.Sprintf("topk(1, sum by (%s) (%s{%s=\"%s\", %s=\"%d\"}))",
		_jsonQueryLabelName, _httpRequestTotalMetricName, _pathLabelName, path, _responseStatusLabelName, responseStatus)
}

func performPrometheusHttpVectorQuery(prometheusQuery string, prometheusVectorQueryURL string) (*prometheusVectorQueryResponse, error) {
	httpRequest, httpRequestErr := http.NewRequest("GET", prometheusVectorQueryURL, nil)
	if httpRequestErr != nil {
		return nil, httpRequestErr
	}

	urlParams := createUrlParamsForPrometheusQuery(prometheusQuery)
	httpRequest.URL.RawQuery = urlParams.Encode()

	httpResponse, httpResponseErr := http.DefaultClient.Do(httpRequest)
	if httpResponseErr != nil {
		return nil, httpResponseErr
	}

	responseBody, httpReaderErr := ioutil.ReadAll(httpResponse.Body)
	if httpReaderErr != nil {
		return nil, httpReaderErr
	}

	vectorResult, unmarshallingErr := parseVectorHttpResponse(responseBody)
	if unmarshallingErr != nil {
		return nil, unmarshallingErr
	}

	return vectorResult, nil
}

func createUrlParamsForPrometheusQuery(prometheusQuery string) url.Values {
	urlParams := url.Values{}
	urlParams.Add(_prometheusVectorQueryParamName, prometheusQuery)
	return urlParams
}

func parseVectorHttpResponse(responseBody []byte) (*prometheusVectorQueryResponse, error) {
	var vectorResult prometheusVectorQueryResponse
	unmarshallingErr := json.Unmarshal(responseBody, &vectorResult)
	if unmarshallingErr != nil {
		return nil, unmarshallingErr
	}

	return &vectorResult, nil
}

func convertVectorResultToMostFrequentQueryRawData(vectorResult *prometheusVectorQueryResponse, target *MostFrequentQueryRawData) error {
	if len(vectorResult.Data.Result) > 0 {
		metric := vectorResult.Data.Result[0].Metric
		if metric != nil && metric[_jsonQueryLabelName] != "" {
			target.RawStrQuery = metric[_jsonQueryLabelName]
		}

		vectorValue := vectorResult.Data.Result[0].Value
		if len(vectorValue) > 1 {
			numberOfHitsStr := vectorValue[1].(string)
			numberOfHitsInt, parsingErr := strconv.Atoi(numberOfHitsStr)
			if parsingErr != nil {
				return parsingErr
			}
			target.NumberOfHits = numberOfHitsInt
		}
	}

	return nil
}
