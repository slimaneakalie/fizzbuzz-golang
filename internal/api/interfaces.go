package api

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

type FizzbuzzRequestAPIHandler interface {
	handleFizzbuzzRequest() fizzhttp.HandlerFunc
}

type StatisticsRequestAPIHandler interface {
	handleStatsRequest() fizzhttp.HandlerFunc
}
