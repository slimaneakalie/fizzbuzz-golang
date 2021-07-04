package fizzhttpMocks

import (
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"
)

func NewTestEngineFactory() *TestingEngineFactory {
	return &TestingEngineFactory{}
}

func (factory *TestingEngineFactory) NewEngine() fizzhttp.Engine {

	return &TestingEngine{}
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
