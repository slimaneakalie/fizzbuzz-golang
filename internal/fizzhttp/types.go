package fizzhttp

import (
	"encoding/json"

	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"
)

type HandlerFunc func(context RequestContext)
type BindingErrorFormatter func(bindError error) error

type requestUnmarshalErrorType = *json.UnmarshalTypeError
type validationErrorsType = validator.ValidationErrors
type bindingBodyType = binding.BindingBody
