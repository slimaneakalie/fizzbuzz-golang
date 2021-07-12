package api

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/logger"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/monitoring"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

func NewRouter(httpEngineFactory fizzhttp.EngineFactory, stringListBuilder stringListBuilder.StringListBuilder,
	monitoringHandler monitoring.Handler, logger logger.Logger) *Router {

	httpEngine := httpEngineFactory.NewEngine()
	httpEngine.GET("/metrics", monitoringHandler.HandleMonitoringQuery())
	httpEngine.UseMiddleware(monitoringHandler.MonitoringMiddleWare())

	fizzbuzzEndpoint := "/v1/fizzbuzz"
	group := httpEngine.Group(fizzbuzzEndpoint)
	fizzbuzzRequestAPIHandler := NewDefaultFizzbuzzRequestAPIHandler(stringListBuilder, httpEngine.FormatBindingError)
	group.POST("", fizzbuzzRequestAPIHandler.handleFizzbuzzRequest())

	statsRequestHandler := NewDefaultStatsHandler(monitoringHandler, fizzbuzzEndpoint, logger)
	group.GET("/stats", statsRequestHandler.handleStatsRequest())

	return &Router{
		HttpEngine: httpEngine,
		logger:     logger,
	}
}

func (router *Router) Run(port int) {
	router.logger.Info("Running server on port", port)
	err := router.HttpEngine.Run(port)
	if err != nil {
		router.logger.Error("Error running server caused by:", err.Error())
	}
}
