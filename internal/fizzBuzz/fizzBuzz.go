package fizzBuzz

func (builder *MainStringListBuilder) BuildStringList(buildInput *StringListBuildInput) []string {
	return []string{buildInput.FirstStr, buildInput.SecondStr}
}
