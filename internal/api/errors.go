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

	case validationErrorType:
		errorObject.Meta = createValidationErrorServerResponse(bindError.(validationErrorType))

	default:
		errorObject.Meta = BadRequestServerResponsePayload{
			Type: BadRequestResponseTypeCode,
		}
	}

	return errorObject
}

func createUnmarshalErrorServerResponse(unmarshalError requestUnmarshalErrorType) BadRequestServerResponsePayload {
	return BadRequestServerResponsePayload{
		Type: BadRequestResponseTypeCode,
		Metadata: []FieldError{{
			Type:   ParamTypeMismatchErrorCode,
			Field:  unmarshalError.Field,
			Detail: generateUnmarshalErrorResponseMessage(unmarshalError),
		}},
	}
}

func generateUnmarshalErrorResponseMessage(unmarshalError requestUnmarshalErrorType) string {
	return fmt.Sprintf("expected type %s instead of %s", unmarshalError.Type.String(), unmarshalError.Value)
}

func createValidationErrorServerResponse(validationError validationErrorType) BadRequestServerResponsePayload {
	var fieldErrors []FieldError
	for _, validationError := range validationError {
		fieldErrors = append(fieldErrors, FieldError{
			Type:   InvalidParamValueErrorCode,
			Field:  validationError.Field(),
			Detail: generateFieldErrorResponseMessage(validationError),
		})
	}

	return BadRequestServerResponsePayload{
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
