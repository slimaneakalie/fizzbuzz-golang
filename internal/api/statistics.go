package api

import (
	"encoding/json"
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/logger"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/monitoring"
)

func NewDefaultStatsHandler(monitoringHelper monitoring.Handler, fizzbuzzEndpoint string, logger logger.Logger) StatisticsRequestAPIHandler {
	return &defaultStatsRequestHandler{
		monitoringHelper: monitoringHelper,
		fizzbuzzEndpoint: fizzbuzzEndpoint,
		logger:           logger,
	}
}

func (handler *defaultStatsRequestHandler) handleStatsRequest() fizzhttp.HandlerFunc {
	return handler.statsRequestHandler
}

func (handler *defaultStatsRequestHandler) statsRequestHandler(context fizzhttp.RequestContext) {
	rawMonitoringData, monitoringErr := handler.monitoringHelper.GetMostFrequentQuery(handler.fizzbuzzEndpoint, http.StatusOK)
	if monitoringErr != nil {
		handler.logger.Error("monitoringHelper.GetMostFrequentQuery", monitoringErr)
		context.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	if rawMonitoringData.RawStrQuery == "{}" {
		context.SendJSONResponse(http.StatusNotFound, nil)
		return
	}

	var mostFrequentRequest FizzbuzzAPIRequest
	unmarshallingErr := json.Unmarshal([]byte(rawMonitoringData.RawStrQuery), &mostFrequentRequest)
	if unmarshallingErr != nil {
		handler.logger.Error("statsRequestHandler.unmarshallingErr", unmarshallingErr)
		context.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	finalStatsAPIResponse := StatsAPIResponse{}
	finalStatsAPIResponse.MostFrequentQuery.Body = mostFrequentRequest
	finalStatsAPIResponse.MostFrequentQuery.NumberOfHits = rawMonitoringData.NumberOfHits

	context.SendJSONResponse(http.StatusOK, finalStatsAPIResponse)
}
