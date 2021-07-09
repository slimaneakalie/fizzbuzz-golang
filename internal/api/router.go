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

	fizzbuzzRequestAPIHandler := NewDefaultFizzbuzzRequestAPIHandler(stringListBuilder, httpEngine.FormatBindingError)
	group := httpEngine.Group("/v1/fizzbuzz")
	group.POST("", fizzbuzzRequestAPIHandler.handleFizzbuzzRequest())

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
