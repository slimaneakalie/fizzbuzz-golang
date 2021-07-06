package common

func (fieldError TestingFieldValidationError) Tag() string {
	return fieldError.TagOutput
}

func (fieldError TestingFieldValidationError) Field() string {
	return fieldError.FieldOutput
}
