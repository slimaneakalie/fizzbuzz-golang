package stringListBuilder

import (
	"strconv"
	"strings"
)

func NewMainStringListBuilder() StringListBuilder {
	return &mainStringListBuilder{}
}

func (builder *mainStringListBuilder) BuildStringList(buildInput *StringListBuildInput) []string {
	stringList := make([]string, buildInput.Limit)
	twoInputStringsJoined := strings.Join([]string{buildInput.FirstStr, buildInput.SecondStr}, "")

	for i := 0; i < buildInput.Limit; i++ {
		currentNumber := i + 1
		firstIntModulo := currentNumber % buildInput.FirstInt
		secondIntModulo := currentNumber % buildInput.SecondInt
		switch {
		case firstIntModulo == 0 && secondIntModulo == 0:
			stringList[i] = twoInputStringsJoined
		case firstIntModulo == 0:
			stringList[i] = buildInput.FirstStr
		case secondIntModulo == 0:
			stringList[i] = buildInput.SecondStr
		default:
			stringList[i] = strconv.Itoa(currentNumber)
		}
	}

	return stringList
}
