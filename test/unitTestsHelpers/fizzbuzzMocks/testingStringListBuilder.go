package fizzbuzzMocks

import "github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"

func NewTestingStringListBuilder() *TestingStringListBuilder {
	return &TestingStringListBuilder{}
}

func (builder *TestingStringListBuilder) BuildStringList(input *stringListBuilder.StringListBuildInput) []string {
	builder.RecordFuncCall("BuildStringList", input)
	return nil
}
