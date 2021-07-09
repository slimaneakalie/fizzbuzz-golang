package apiTestTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzbuzzMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzhttpMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/loggerMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/monitoringMocks"
)

func PrepareRouterTestsTooling(engineRunMethodOutput error) *Router {
	testingEngine := fizzhttpMocks.NewTestingEngine(engineRunMethodOutput, nil)

	return &Router{
		TestingEngine:            testingEngine,
		TestingEngineFactory:     fizzhttpMocks.NewTestingEngineFactory(testingEngine),
		TestingStringListBuilder: fizzbuzzMocks.NewTestingStringListBuilder(nil),
		TestingLogger:            loggerMocks.NewTestingLogger(),
		TestingMonitoringHandler: monitoringMocks.NewTestingHandler(),
	}
}
