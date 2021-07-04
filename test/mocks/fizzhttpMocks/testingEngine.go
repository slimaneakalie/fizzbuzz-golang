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
	factory.RecordFuncCall("NewEngine", []interface{}{})
	return factory.engine
}

func (engine *TestingEngine) Group(relativePath string) fizzhttp.RouterGroup {
	return &TestingRouterGroup{}
}

func (engine *TestingEngine) FormatBindingError(bindError error) error {
	return nil
}

func (engine *TestingEngine) Run(addr string) error {
	return nil
}
