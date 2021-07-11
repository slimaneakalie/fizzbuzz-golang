package monitoring

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func NewPrometheusHandler() Handler {
	return &PrometheusHandler{
		promHttpHandler: promhttp.Handler(),
		httpRequestsCounterVector: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "The total number of http requests",
		}, []string{"path", "json_query", "response_status"}),
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

func (handler *PrometheusHandler) GetMostFrequentQuery() MostFrequentQueryRawData {
	// MOST_FREQUENT_QUERY= topk(1, sum by (json_query) (http_request_total{path="/v1/fizzbuzz", response_status="200"}))
	// PROMETHEUS_ENDPOINT = /api/v1/query
	// PROMETHEUS_QS_PARAM_NAME = query
	// PROMETHEUS_HOST -> should be a field name that we get from a config = http://localhost:9090
	// Endpoint format:= ${PROMETHEUS_HOST}/${PROMETHEUS_ENDPOINT}?PROMETHEUS_QS_PARAM_NAME=MOST_FREQUENT_QUERY -> Use http helpers to encode the query
	// Deserialize the promotheus response
	// Handle the case where data.result is nil ou have length < 1
	// Use strconv.Atoi(data.result[0].value[1]) on NumberOfHits
	// Use data.result[0].metric.json_query on RawStrQuery
}
