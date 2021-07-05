package apiTestTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzbuzzMocks"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/fizzhttpMocks"
)

func PrepareFizzbuzzTestsTooling(buildStringListOutput []string, shouldBindBodyWithOutput error, formatBindingErrorOutput error) *Fizzbuzz {
	return &Fizzbuzz{
		TestingEngine:            fizzhttpMocks.NewTestingEngine(nil, formatBindingErrorOutput),
		TestingRequestContext:    fizzhttpMocks.NewTestingRequestContext(shouldBindBodyWithOutput),
		TestingStringListBuilder: fizzbuzzMocks.NewTestingStringListBuilder(buildStringListOutput),
	}
}
