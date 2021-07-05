package apiTestTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzbuzzMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzhttpMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/loggerMocks"
)

type Router struct {
	TestingEngine            *fizzhttpMocks.TestingEngine
	TestingEngineFactory     *fizzhttpMocks.TestingEngineFactory
	TestingStringListBuilder *fizzbuzzMocks.TestingStringListBuilder
	TestingLogger            *loggerMocks.TestingLogger
}
