package fizzbuzzMocks

import "github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"

func NewTestingStringListBuilder(buildStringListOutput []string) *TestingStringListBuilder {
	return &TestingStringListBuilder{
		buildStringListOutput: buildStringListOutput,
	}
}

func (builder *TestingStringListBuilder) BuildStringList(input *stringListBuilder.StringListBuildInput) []string {
	builder.RecordFuncCall("BuildStringList", input)
	return builder.buildStringListOutput
}
