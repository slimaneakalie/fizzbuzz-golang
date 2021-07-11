package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusHandler struct {
	promHttpHandler           http.Handler
	httpRequestsCounterVector *prometheus.CounterVec
}

type MostFrequentQueryRawData struct {
	RawStrQuery  string
	NumberOfHits int
}
