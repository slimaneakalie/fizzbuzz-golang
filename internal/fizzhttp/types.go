package fizzhttp

import (
	"encoding/json"
)

type HandlerFunc func(context RequestContext)
type BindingErrorFormatter func(bindError error) error

type requestUnmarshalErrorType = *json.UnmarshalTypeError
type validationErrorsType []fieldValidationError
