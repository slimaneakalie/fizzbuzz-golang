package main

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/service"
)

func main() {
	port := 9000                                   // TODO get this from a config
	prometheusHostName := "http://prometheus:9090" // TODO get this from a config
	service.Start(port, prometheusHostName)
}
