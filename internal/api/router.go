package api

import (
	"fmt"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzbuzz"
)

func NewRouter(httpEngineFactory fizzhttp.EngineFactory, stringListBuilder fizzbuzz.StringListBuilder) *Router {
	engine := httpEngineFactory.NewEngine()

	fizzbuzzRequestAPIHandler := MainFizzbuzzRequestAPIHandler{
		StringListBuilder:     stringListBuilder,
		bindingErrorFormatter: engine.FormatBindingError,
	}

	engine.Group("/v1/fizzbuzz").
		POST("/", fizzbuzzRequestAPIHandler.handleFizzbuzzRequest())

	return &Router{
		httpEngine: engine,
	}
}

func (router *Router) Run(port int) {
	fmt.Println("Running server on port ", port)
	err := router.httpEngine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error running server caused by: ", err.Error())
	}
}
