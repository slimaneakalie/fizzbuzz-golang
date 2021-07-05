package main

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/logger"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

func main() {
	serviceMode := fizzhttp.ReleaseMode // TODO get this from a config
	mainEngineFactory := fizzhttp.NewEngineFactory(serviceMode)

	mainBuilder := stringListBuilder.NewMainStringListBuilder()
	mainLogger := logger.NewMainLogger()

	router := api.NewRouter(mainEngineFactory, mainBuilder, mainLogger)

	port := 9000 // TODO get this from a config
	router.Run(port)
}
