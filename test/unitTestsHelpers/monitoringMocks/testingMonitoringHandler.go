package monitoringMocks

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/monitoring"
)

func NewTestingHandler() *TestingHandler {
	return &TestingHandler{}
}

func (handler *TestingHandler) HandleMonitoringQuery() fizzhttp.HandlerFunc {
	handler.RecordFuncCall("HandleMonitoringQuery")
	return func(context fizzhttp.RequestContext) {}
}

func (handler *TestingHandler) MonitoringMiddleWare() fizzhttp.HandlerFunc {
	handler.RecordFuncCall("MonitoringMiddleWare")
	return func(context fizzhttp.RequestContext) {}
}

func (handler *TestingHandler) GetMostFrequentQuery(path string, responseStatus int) (*monitoring.MostFrequentQueryRawData, error) {
	handler.RecordFuncCall("GetMostFrequentQuery", path, responseStatus)
	return &monitoring.MostFrequentQueryRawData{}, nil
}
