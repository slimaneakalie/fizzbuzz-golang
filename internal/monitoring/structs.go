package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusHandler struct {
	promHttpHandler           http.Handler
	httpRequestsCounterVector *prometheus.CounterVec
	prometheusVectorQueryURL  string
}

type MostFrequentQueryRawData struct {
	RawStrQuery  string
	NumberOfHits int
}

type prometheusVectorQueryResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
