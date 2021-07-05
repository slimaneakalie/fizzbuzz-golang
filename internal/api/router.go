package api

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/logger"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

func NewRouter(httpEngineFactory fizzhttp.EngineFactory, stringListBuilder stringListBuilder.StringListBuilder, logger logger.Logger) *router {
	httpEngine := httpEngineFactory.NewEngine()

	fizzbuzzRequestAPIHandler := NewMainFizzbuzzRequestAPIHandler(stringListBuilder, httpEngine.FormatBindingError)

	group := httpEngine.Group("/v1/stringListBuilder")
	group.POST("/", fizzbuzzRequestAPIHandler.handleFizzbuzzRequest())

	return &router{
		httpEngine: httpEngine,
		logger:     logger,
	}
}

func (router *router) Run(port int) {
	router.logger.Info("Running server on port", port)
	err := router.httpEngine.Run(port)
	if err != nil {
		router.logger.Error("Error running server caused by:", err.Error())
	}
}
