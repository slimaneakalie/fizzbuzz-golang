package api

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

type FizzbuzzRequestAPIHandler interface {
	handleFizzbuzzRequest() fizzhttp.HandlerFunc
}
