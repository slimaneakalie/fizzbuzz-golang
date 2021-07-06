package fizzhttp

func NewHttpErrorResponseMetadata(errorType string, fieldErrors []*fieldError) *httpErrorResponseMetadata {
	return &httpErrorResponseMetadata{
		Type:        errorType,
		FieldErrors: fieldErrors,
	}
}

func NewFieldError(errorType, field, detail string) *fieldError {
	return &fieldError{
		Type:   errorType,
		Field:  field,
		Detail: detail,
	}
}
