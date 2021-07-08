package fizzhttp

func NewHttpErrorResponse(errorType string, fieldErrors []*responseFieldError) *httpErrorResponse {
	return &httpErrorResponse{
		Type:        errorType,
		FieldErrors: fieldErrors,
	}
}

func NewFieldError(errorType, field, detail string) *responseFieldError {
	return &responseFieldError{
		Type:   errorType,
		Field:  field,
		Detail: detail,
	}
}
