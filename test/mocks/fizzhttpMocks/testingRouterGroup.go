package fizzhttpMocks

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func (group *TestingRouterGroup) POST(relativePath string, handler fizzhttp.HandlerFunc) {
}

func (group *TestingRouterGroup) GET(relativePath string, handler fizzhttp.HandlerFunc) {
}
