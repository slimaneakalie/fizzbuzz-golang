package fizzhttpMocks

import (
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func NewTestingEngineFactory(testingEngine *TestingEngine) *TestingEngineFactory {
	return &TestingEngineFactory{
		engine: testingEngine,
	}
}

func NewTestingEngine(runMethodOutput error, formatBindingErrorOutput error) *TestingEngine {
	return &TestingEngine{
		RouterGroups:             []*TestingRouterGroup{},
		runMethodOutput:          runMethodOutput,
		formatBindingErrorOutput: formatBindingErrorOutput,
	}
}

func (factory *TestingEngineFactory) NewEngine() fizzhttp.Engine {
	factory.RecordFuncCall("NewEngine")
	return factory.engine
}

func (engine *TestingEngine) Group(relativePath string) fizzhttp.RouterGroup {
	engine.RecordFuncCall("Group", relativePath)
	newRouterGroup := &TestingRouterGroup{}
	engine.RouterGroups = append(engine.RouterGroups, newRouterGroup)
	return newRouterGroup
}

func (engine *TestingEngine) FormatBindingError(bindError error) error {
	engine.RecordFuncCall("FormatBindingError", bindError)
	return engine.formatBindingErrorOutput
}

func (engine *TestingEngine) Run(port int) error {
	engine.RecordFuncCall("Run", port)
	return engine.runMethodOutput
}

func (engine *TestingEngine) ServeHTTP(respWriter http.ResponseWriter, request *http.Request) {
	engine.RecordFuncCall("ServeHTTP", respWriter, request)
}

func (engine *TestingEngine) POST(relativePath string, handler fizzhttp.HandlerFunc) {
	engine.RecordFuncCall("POST", relativePath, handler)
}

func (engine *TestingEngine) GET(relativePath string, handler fizzhttp.HandlerFunc) {
	engine.RecordFuncCall("GET", relativePath, handler)
}

func (engine *TestingEngine) UseMiddleware(handlerFunc fizzhttp.HandlerFunc) {
	engine.RecordFuncCall("UseMiddleware", handlerFunc)
}
