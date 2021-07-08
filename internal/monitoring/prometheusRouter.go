package monitoring

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func NewPrometheusHandler() Handler {
	return &PrometheusHandler{
		promHttpHandler: promhttp.Handler(),
	}
}

func (handler *PrometheusHandler) HandleMonitoringQuery() fizzhttp.HandlerFunc {
	return handler.monitoringQueryHandler
}

func (handler *PrometheusHandler) monitoringQueryHandler(context fizzhttp.RequestContext) {
	handler.promHttpHandler.ServeHTTP(context.GetResponseWriter(), context.GetRequest())
}
