package fizzbuzzMocks

import "github.com/slimaneakalie/fizzbuzz-golang/internal/fizzbuzz"

func NewTestingStringListBuilder() *TestingStringListBuilder {
	return &TestingStringListBuilder{}
}

func (builder *TestingStringListBuilder) BuildStringList(input *fizzbuzz.StringListBuildInput) []string {
	return nil
}
