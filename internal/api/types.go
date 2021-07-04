package api

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzBuzz"

	"github.com/go-playground/validator/v10"
)

type Router struct {
	ginEngine *gin.Engine
}

type FizzBuzzRequestAPIHandler interface {
	handleFizzBuzzRequest() gin.HandlerFunc
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

type badRequestServerResponsePayload struct {
	Type     string       `json:"type"`
	Metadata []fieldError `json:"errors,omitempty"`
}

type fieldError struct {
	Type   string `json:"type"`
	Field  string `json:"field"`
	Detail string `json:"detail"`
}

type requestUnmarshalErrorType = *json.UnmarshalTypeError
type validationErrorsType = validator.ValidationErrors
