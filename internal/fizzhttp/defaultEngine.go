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

func (engine *defaultEngine) FormatBindingError(bindingError error) error {
	formattedError := &gin.Error{
		Err:  bindingError,
		Type: gin.ErrorTypePublic,
	}

	switch bindingError.(type) {
	case requestUnmarshalErrorType:
		formattedError.Meta = createUnmarshalErrorServerResponse(bindingError.(requestUnmarshalErrorType))

	case validationErrorsType:
		formattedError.Meta = createValidationErrorServerResponse(bindingError.(validationErrorsType))

	default:
		formattedError.Meta = httpErrorResponseMetadata{
			Type: BadRequestResponseTypeCode,
		}
	}

	return formattedError
}
