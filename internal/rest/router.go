package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzBuzz"
)

func NewRouter(serviceMode string, stringListBuilder fizzBuzz.StringListBuilder) *Router {
	gin.SetMode(serviceMode)
	fizzBuzzRequestAPIHandler := api.MainFizzBuzzRequestAPIHandler{
		StringListBuilder: stringListBuilder,
	}

	engine := gin.Default()
	engine.Group("/v1/fizzbuzz").
		POST("/", fizzBuzzRequestAPIHandler.HandleFizzBuzzRequest())

	return &Router{
		ginEngine: engine,
	}
}

func (router *Router) Run(port int) {
	fmt.Println("Running server on port ", port)

	err := router.ginEngine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error running server caused by: ", err.Error())
	}
}
