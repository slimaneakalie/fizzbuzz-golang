package service

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/logger"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/monitoring"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

func Start(port int, prometheusHostName string) {
	monitoringHandler := monitoring.NewPrometheusHandler(prometheusHostName)
	server := NewServer(monitoringHandler)

	server.Run(port)
}

func NewServer(monitoringHandler monitoring.Handler) *api.Router {
	serviceMode := fizzhttp.ReleaseMode
	httpEngineFactory := fizzhttp.NewDefaultHttpEngineFactory(serviceMode)

	stringListBuilder := stringListBuilder.NewDefaultBuilder()
	logger := logger.NewDefaultLogger()

	router := api.NewRouter(httpEngineFactory, stringListBuilder, monitoringHandler, logger)
	return router
}
