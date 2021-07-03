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
	FirstInt  int    `json:"firstInt" binding:"required"`
	SecondInt int    `json:"secondInt" binding:"required"`
	Limit     int    `json:"limit" binding:"required"`
	FirstStr  string `json:"firstStr" binding:"required"`
	SecondStr string `json:"secondStr" binding:"required"`
}

type FizzBuzzAPIResponse struct {
	FizzBuzzStringList []string `json:"fizzBuzzStringList"`
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
