package apiTestTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzbuzzMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzhttpMocks"
)

type Fizzbuzz struct {
	TestingEngine            *fizzhttpMocks.TestingEngine
	TestingRequestContext    *fizzhttpMocks.TestingRequestContext
	TestingStringListBuilder *fizzbuzzMocks.TestingStringListBuilder
}
