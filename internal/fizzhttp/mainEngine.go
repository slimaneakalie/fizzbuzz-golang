package fizzhttp

import "github.com/gin-gonic/gin"

func NewEngineFactory(serviceMode string) *MainEngineFactory {
	gin.SetMode(serviceMode)
	return &MainEngineFactory{}
}

func (factory *MainEngineFactory) NewEngine() Engine {
	return &mainEngine{
		internalEngine: gin.Default(),
	}
}

func (engine *mainEngine) Group(relativePath string) RouterGroup {
	return &mainRouterGroup{
		internalRouterGroup: engine.internalEngine.Group(relativePath),
	}
}

func (engine *mainEngine) Run(addr string) error {
	return engine.internalEngine.Run(addr)
}

func (engine *mainEngine) FormatBindingError(bindError error) error {
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