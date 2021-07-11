package api

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/logger"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/monitoring"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

type Router struct {
	HttpEngine fizzhttp.Engine
	logger     logger.Logger
}

type defaultFizzbuzzRequestAPIHandler struct {
	StringListBuilder     stringListBuilder.StringListBuilder
	bindingErrorFormatter fizzhttp.BindingErrorFormatter
}

type defaultStatsRequestHandler struct {
	monitoringHelper monitoring.Handler
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

type StatsAPIResponse struct {
	MostFrequentQuery struct {
		Body         FizzbuzzAPIRequest `json:"body"`
		NumberOfHits int                `json:"numberOfHits"`
	} `json:"mostFrequentQuery"`
}
