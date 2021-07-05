package fizzhttp

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewDefaultHttpEngineFactory(serviceMode string) EngineFactory {
	gin.SetMode(serviceMode)
	return &defaultEngineFactory{}
}

func (factory *defaultEngineFactory) NewEngine() Engine {
	return &defaultEngine{
		internalEngine: gin.Default(),
	}
}

func (engine *defaultEngine) Group(relativePath string) RouterGroup {
	return &defaultRouterGroup{
		internalRouterGroup: engine.internalEngine.Group(relativePath),
	}
}

func (engine *defaultEngine) Run(port int) error {
	return engine.internalEngine.Run(fmt.Sprintf(":%d", port))
}

func (engine *defaultEngine) FormatBindingError(bindError error) error {
	formattedError := &gin.Error{
		Err: bindError,
	}

	switch bindError.(type) {
	case requestUnmarshalErrorType:
		formattedError.Meta = createUnmarshalErrorServerResponse(bindError.(requestUnmarshalErrorType))

	case validationErrorsType:
		formattedError.Meta = createValidationErrorServerResponse(bindError.(validationErrorsType))

	default:
		formattedError.Meta = badRequestServerResponsePayload{
			Type: BadRequestResponseTypeCode,
		}
	}

	return formattedError
}
