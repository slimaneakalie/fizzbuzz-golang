package api

import (
	"encoding/json"
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/monitoring"
)

func NewDefaultStatsHandler(monitoringHelper monitoring.Handler) StatisticsRequestAPIHandler {
	return &defaultStatsRequestHandler{
		monitoringHelper: monitoringHelper,
	}
}

func (handler *defaultStatsRequestHandler) handleStatsRequest() fizzhttp.HandlerFunc {
	return handler.statsRequestHandler
}

func (handler *defaultStatsRequestHandler) statsRequestHandler(context fizzhttp.RequestContext) {
	rawMonitoringData := handler.monitoringHelper.GetMostFrequentQuery()

	var mostFrequentRequest FizzbuzzAPIRequest
	unmarshallingErr := json.Unmarshal([]byte(rawMonitoringData.RawStrQuery), &mostFrequentRequest)
	if unmarshallingErr != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	finalStatsAPIResponse := StatsAPIResponse{}
	finalStatsAPIResponse.MostFrequentQuery.Body = mostFrequentRequest
	finalStatsAPIResponse.MostFrequentQuery.NumberOfHits = rawMonitoringData.NumberOfHits

	context.SendJSONResponse(http.StatusOK, finalStatsAPIResponse)
}
