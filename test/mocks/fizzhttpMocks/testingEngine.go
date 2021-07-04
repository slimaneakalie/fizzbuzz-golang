package fizzhttpMocks

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func NewTestingEngineFactory(testingEngine *TestingEngine) *TestingEngineFactory {
	return &TestingEngineFactory{
		engine: testingEngine,
	}
}

func (factory *TestingEngineFactory) NewEngine() fizzhttp.Engine {
	factory.RecordFuncCall("NewEngine")
	return factory.engine
}

func (engine *TestingEngine) Group(relativePath string) fizzhttp.RouterGroup {
	engine.RecordFuncCall("Group", relativePath)
	return &TestingRouterGroup{}
}

func (engine *TestingEngine) FormatBindingError(bindError error) error {
	engine.RecordFuncCall("FormatBindingError", bindError)
	return nil
}

func (engine *TestingEngine) Run(addr string) error {
	engine.RecordFuncCall("Run", addr)
	return nil
}
