package fizzhttpMocks

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func (group *TestingRouterGroup) POST(relativePath string, handler fizzhttp.HandlerFunc) {
	group.RecordFuncCall("POST", relativePath, handler)
}

func (group *TestingRouterGroup) GET(relativePath string, handler fizzhttp.HandlerFunc) {
	group.RecordFuncCall("GET", relativePath, handler)
}
