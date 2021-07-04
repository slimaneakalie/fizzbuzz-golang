package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	BadRequestResponseTypeCode = "BadRequest"
	InvalidParamValueErrorCode = "InvalidParamValue"
	ParamTypeMismatchErrorCode = "ParamTypeMismatch"
)

func handleBindError(bindError error) *gin.Error {
	errorObject := &gin.Error{
		Err: bindError,
	}

	switch bindError.(type) {
	case requestUnmarshalErrorType:
		errorObject.Meta = createUnmarshalErrorServerResponse(bindError.(requestUnmarshalErrorType))

	case validationErrorsType:
		errorObject.Meta = createValidationErrorServerResponse(bindError.(validationErrorsType))

	default:
		errorObject.Meta = badRequestServerResponsePayload{
			Type: BadRequestResponseTypeCode,
		}
	}

	return errorObject
}

func createUnmarshalErrorServerResponse(unmarshalError requestUnmarshalErrorType) badRequestServerResponsePayload {
	return badRequestServerResponsePayload{
		Type: BadRequestResponseTypeCode,
		Metadata: []fieldError{{
			Type:   ParamTypeMismatchErrorCode,
			Field:  unmarshalError.Field,
			Detail: generateUnmarshalErrorResponseMessage(unmarshalError),
		}},
	}
}

func generateUnmarshalErrorResponseMessage(unmarshalError requestUnmarshalErrorType) string {
	return fmt.Sprintf("expected type %s instead of %s", unmarshalError.Type.String(), unmarshalError.Value)
}

func createValidationErrorServerResponse(validationErrors validationErrorsType) badRequestServerResponsePayload {
	var fieldErrors []fieldError
	for _, validationError := range validationErrors {
		fieldErrors = append(fieldErrors, fieldError{
			Type:   InvalidParamValueErrorCode,
			Field:  validationError.Field(),
			Detail: generateFieldErrorResponseMessage(validationError),
		})
	}

	return badRequestServerResponsePayload{
		Type:     BadRequestResponseTypeCode,
		Metadata: fieldErrors,
	}
}

func generateFieldErrorResponseMessage(ve validator.FieldError) string {
	switch ve.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", ve.Field())

	default:
		return fmt.Sprintf("Invalid value for field %s ", ve.Field())
	}
}
