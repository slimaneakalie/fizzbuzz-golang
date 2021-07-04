package testTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzbuzzMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzhttpMocks"
)

type RouterTestsTooling struct {
	TestingEngine        *fizzhttpMocks.TestingEngine
	TestingEngineFactory *fizzhttpMocks.TestingEngineFactory
	StringListBuilder    *fizzbuzzMocks.TestingStringListBuilder
}
