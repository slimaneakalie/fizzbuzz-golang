package monitoring

import "github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"

type Handler interface {
	HandleMonitoringQuery() fizzhttp.HandlerFunc
	MonitoringMiddleWare() fizzhttp.HandlerFunc
	GetMostFrequentQuery() MostFrequentQueryRawData
}
