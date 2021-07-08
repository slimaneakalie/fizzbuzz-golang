package fizzhttp

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	BadRequestResponseTypeCode = "BadRequest"
	InvalidParamValueErrorCode = "InvalidParamValue"
	ParamTypeMismatchErrorCode = "ParamTypeMismatch"
)

func createUnmarshalErrorServerResponse(unmarshalError requestUnmarshalErrorType) *httpErrorResponse {
	errorDetail := generateUnmarshalErrorResponseMessage(unmarshalError)
	currentFieldError := NewFieldError(ParamTypeMismatchErrorCode, unmarshalError.Field, errorDetail)
	responseFieldErrors := []*responseFieldError{currentFieldError}

	return NewHttpErrorResponse(BadRequestResponseTypeCode, responseFieldErrors)
}

func generateUnmarshalErrorResponseMessage(unmarshalError requestUnmarshalErrorType) string {
	return fmt.Sprintf("expected type %s instead of %s", unmarshalError.Type.String(), unmarshalError.Value)
}

func createValidationErrorServerResponse(validationErrors validationErrorsType) *httpErrorResponse {
	var responseFieldErrors []*responseFieldError
	for _, validationError := range validationErrors {
		currentFieldErrorDetail := generateFieldErrorResponseMessage(validationError)
		currentFieldError := NewFieldError(InvalidParamValueErrorCode, validationError.Field(), currentFieldErrorDetail)
		responseFieldErrors = append(responseFieldErrors, currentFieldError)
	}

	return NewHttpErrorResponse(BadRequestResponseTypeCode, responseFieldErrors)
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

func (err *httpErrorResponse) Error() string {
	jsonBytes, _ := json.Marshal(err)
	return string(jsonBytes)
}
