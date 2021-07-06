package fizzhttp

import (
	"fmt"
)

const (
	BadRequestResponseTypeCode = "BadRequest"
	InvalidParamValueErrorCode = "InvalidParamValue"
	ParamTypeMismatchErrorCode = "ParamTypeMismatch"
)

func createUnmarshalErrorServerResponse(unmarshalError requestUnmarshalErrorType) *httpErrorResponseMetadata {
	errorDetail := generateUnmarshalErrorResponseMessage(unmarshalError)
	currentFieldError := NewFieldError(ParamTypeMismatchErrorCode, unmarshalError.Field, errorDetail)
	responseFieldErrors := []*fieldError{currentFieldError}

	return NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, responseFieldErrors)
}

func generateUnmarshalErrorResponseMessage(unmarshalError requestUnmarshalErrorType) string {
	return fmt.Sprintf("expected type %s instead of %s", unmarshalError.Value, unmarshalError.Type.String())
}

func createValidationErrorServerResponse(validationErrors validationErrorsType) *httpErrorResponseMetadata {
	var responseFieldErrors []*fieldError
	for _, validationError := range validationErrors {
		currentFieldErrorDetail := generateFieldErrorResponseMessage(validationError)
		currentFieldError := NewFieldError(InvalidParamValueErrorCode, validationError.Field(), currentFieldErrorDetail)
		responseFieldErrors = append(responseFieldErrors, currentFieldError)
	}

	return NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, responseFieldErrors)
}

func generateFieldErrorResponseMessage(ve fieldValidationError) string {
	switch ve.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", ve.Field())

	default:
		return fmt.Sprintf("Invalid value for field %s ", ve.Field())
	}
}
