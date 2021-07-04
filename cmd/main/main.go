package main

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzbuzz"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func main() {
	serviceMode := fizzhttp.ReleaseMode // TODO get this from a config
	mainEngineFactory := fizzhttp.NewEngineFactory(serviceMode)
	mainBuilder := &fizzbuzz.MainStringListBuilder{}
	router := api.NewRouter(mainEngineFactory, mainBuilder)

	port := 9000 // TODO get this from a config
	router.Run(port)
}
