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
