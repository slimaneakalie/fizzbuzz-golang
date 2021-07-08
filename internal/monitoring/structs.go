package monitoring

import "net/http"

type PrometheusHandler struct {
	promHttpHandler http.Handler
}
