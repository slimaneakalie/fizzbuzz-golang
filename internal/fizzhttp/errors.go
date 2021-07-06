package fizzhttp

import (
	"fmt"
	"strings"
)

const (
	BadRequestResponseTypeCode = "BadRequest"
	InvalidParamValueErrorCode = "InvalidParamValue"
	ParamTypeMismatchErrorCode = "ParamTypeMismatch"
)

func createUnmarshalErrorServerResponse(unmarshalError requestUnmarshalErrorType) *httpErrorResponseMetadata {
	errorDetail := generateUnmarshalErrorResponseMessage(unmarshalError)
	currentFieldError := NewFieldError(ParamTypeMismatchErrorCode, unmarshalError.Field, errorDetail)
	responseFieldErrors := []*responseFieldError{currentFieldError}

	return NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, responseFieldErrors)
}

func generateUnmarshalErrorResponseMessage(unmarshalError requestUnmarshalErrorType) string {
	return fmt.Sprintf("expected type %s instead of %s", unmarshalError.Value, unmarshalError.Type.String())
}

func createValidationErrorServerResponse(validationErrors validationErrorsType) *httpErrorResponseMetadata {
	var responseFieldErrors []*responseFieldError
	for _, validationError := range validationErrors {
		currentFieldErrorDetail := generateFieldErrorResponseMessage(validationError)
		currentFieldError := NewFieldError(InvalidParamValueErrorCode, validationError.Field(), currentFieldErrorDetail)
		responseFieldErrors = append(responseFieldErrors, currentFieldError)
	}

	return NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, responseFieldErrors)
}

func generateFieldErrorResponseMessage(fve fieldValidationError) string {
	switch fve.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fve.Field())

	default:
		return fmt.Sprintf("Invalid value for field %s", fve.Field())
	}
}

func (validationErrors validationErrorsType) Error() string {
	var errorMessages []string
	for _, vError := range validationErrors {
		errorMessage := generateFieldErrorResponseMessage(vError)
		errorMessages = append(errorMessages, errorMessage)
	}

	return strings.Join(errorMessages, ",\n")
}
