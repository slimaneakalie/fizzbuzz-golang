package testTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzbuzzMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzhttpMocks"
)

func PrepareRouterTestsTooling(engineRunMethodOutput error) *RouterTestsTooling {
	testingEngine := fizzhttpMocks.NewTestingEngine(engineRunMethodOutput)

	return &RouterTestsTooling{
		TestingEngine:        testingEngine,
		TestingEngineFactory: fizzhttpMocks.NewTestingEngineFactory(testingEngine),
		StringListBuilder:    fizzbuzzMocks.NewTestingStringListBuilder(),
	}
}
