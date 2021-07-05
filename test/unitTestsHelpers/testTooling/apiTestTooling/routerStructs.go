package apiTestTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzbuzzMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzhttpMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/loggerMocks"
)

type RouterTestsTooling struct {
	TestingEngine        *fizzhttpMocks.TestingEngine
	TestingEngineFactory *fizzhttpMocks.TestingEngineFactory
	StringListBuilder    *fizzbuzzMocks.TestingStringListBuilder
	TestingLogger        *loggerMocks.TestingLogger
}
