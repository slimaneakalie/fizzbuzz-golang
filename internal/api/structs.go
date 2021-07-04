package api

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzbuzz"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

type Router struct {
	httpEngine fizzhttp.Engine
}

type MainFizzbuzzRequestAPIHandler struct {
	StringListBuilder     fizzbuzz.StringListBuilder
	bindingErrorFormatter fizzhttp.BindingErrorFormatter
}

type FizzbuzzAPIRequest struct {
	FirstInt  int    `json:"firstInt" binding:"required"`
	SecondInt int    `json:"secondInt" binding:"required"`
	Limit     int    `json:"limit" binding:"required"`
	FirstStr  string `json:"firstStr" binding:"required"`
	SecondStr string `json:"secondStr" binding:"required"`
}

type FizzbuzzAPIResponse struct {
	FizzbuzzStringList []string `json:"fizzbuzzStringList"`
}
