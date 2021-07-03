package api

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzBuzz"

	"github.com/go-playground/validator/v10"
)

type FizzBuzzRequestAPIHandler interface {
	HandleFizzBuzzRequest() gin.HandlerFunc
}

type MainFizzBuzzRequestAPIHandler struct {
	StringListBuilder fizzBuzz.StringListBuilder
}

type FizzBuzzAPIRequest struct {
	firstInt  int    `json:"firstInt": "required"`
	secondInt int    `json:"secondInt": "required"`
	thirdInt  int    `json:"thirdInt": "required"`
	limit     int    `json:"limit": "required"`
	firstStr  string `json:"firstStr": "required"`
	secondStr string `json:"secondStr": "required"`
}

type FizzBuzzAPIResponse struct {
	fizzBuzzStringList []string `json:"fizzBuzzStringList"`
}

type BadRequestServerResponsePayload struct {
	Type     string       `json:"type"`
	Metadata []FieldError `json:"errors,omitempty"`
}

type FieldError struct {
	Type   string `json:"type"`
	Field  string `json:"field"`
	Detail string `json:"detail"`
}

type requestUnmarshalErrorType = *json.UnmarshalTypeError
type validationErrorType = validator.ValidationErrors
