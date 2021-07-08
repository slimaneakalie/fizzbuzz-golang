package monitoring

import "github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"

type Router interface {
	HandleMonitoringQuery() fizzhttp.HandlerFunc
}
