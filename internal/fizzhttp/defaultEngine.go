package fizzhttp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

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

	switch bindingError.(type) {
	case *json.UnmarshalTypeError:
		return createUnmarshalErrorServerResponse(bindingError.(*json.UnmarshalTypeError))

	case validationErrorsType:
		fieldValidationErrors := bindingError.(validationErrorsType)
		return createValidationErrorServerResponse(fieldValidationErrors)

	case validator.ValidationErrors:
		validatorErrors := bindingError.(validator.ValidationErrors)
		var fieldValidationErrors validationErrorsType
		for _, ve := range validatorErrors {
			fieldValidationErrors = append(fieldValidationErrors, ve)
		}

		return createValidationErrorServerResponse(fieldValidationErrors)

	default:
		return NewHttpErrorResponse(BadRequestResponseTypeCode, nil)
	}

}

func (engine *defaultEngine) ServeHTTP(respWriter http.ResponseWriter, request *http.Request) {
	engine.internalEngine.ServeHTTP(respWriter, request)
}
