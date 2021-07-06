package fizzhttp

import (
	"encoding/json"

	"github.com/gin-gonic/gin/binding"
)

type HandlerFunc func(context RequestContext)
type BindingErrorFormatter func(bindError error) error

type requestUnmarshalErrorType = *json.UnmarshalTypeError
type bindingBodyType = binding.BindingBody
type validationErrorsType []fieldValidationError
