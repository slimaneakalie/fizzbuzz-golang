package fizzbuzzMocks

import "github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/common"

type TestingStringListBuilder struct {
	common.MockElement
	buildStringListOutput []string
}
