package main

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/logger"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

func main() {
	serviceMode := fizzhttp.ReleaseMode // TODO get this from a config
	httpEngineFactory := fizzhttp.NewDefaultHttpEngineFactory(serviceMode)

	stringListBuilder := stringListBuilder.NewDefaultStringListBuilder()
	logger := logger.NewDefaultLogger()

	router := api.NewRouter(httpEngineFactory, stringListBuilder, logger)

	port := 9000 // TODO get this from a config
	router.Run(port)
}
