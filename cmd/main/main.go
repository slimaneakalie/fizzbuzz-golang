package main

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzBuzz"

	"github.com/gin-gonic/gin"
)

func main() {
	serviceMode := gin.ReleaseMode // TODO get this from a config
	mainBuilder := &fizzBuzz.MainStringListBuilder{}
	router := api.NewRouter(serviceMode, mainBuilder)

	port := 9000 // TODO get this from a config
	router.Run(port)
}
