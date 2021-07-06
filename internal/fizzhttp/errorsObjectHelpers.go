package fizzhttp

func NewHttpErrorResponseMetadata(errorType string, fieldErrors []*responseFieldError) *httpErrorResponseMetadata {
	return &httpErrorResponseMetadata{
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
