package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
)

func main() {
	router := gin.Default()
	apiVersion := "v1"
	router.Group(fmt.Sprintf("/%s/search", apiVersion)).
		POST("/", api.HandleFizzBuzzRequest())
}
