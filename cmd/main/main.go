package main

import (
	"fmt"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzBuzz"

	"github.com/gin-gonic/gin"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
)

func main() {
	// TODO separate this into dedicated files and functions
	serviceMode := gin.ReleaseMode // TODO get this from a config
	gin.SetMode(serviceMode)

	mainBuilder := &fizzBuzz.MainStringListBuilder{}

	fizzBuzzRequestAPIHandler := api.MainFizzBuzzRequestAPIHandler{
		StringListBuilder: mainBuilder,
	}

	router := gin.Default()
	port := 9000 // TODO get this from a config

	router.Group("/v1/fizzbuzz").
		POST("/", fizzBuzzRequestAPIHandler.HandleFizzBuzzRequest())

	fmt.Printf("Running server on %d\n", port)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error running server caused by: %s\n", err.Error())
	}
}
