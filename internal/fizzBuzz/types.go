package fizzBuzz

type StringListBuildInput struct {
	FirstInt  int
	SecondInt int
	ThirdInt  int
	Limit     int
	FirstStr  string
	SecondStr string
}

type StringListBuilder interface {
	BuildStringList(*StringListBuildInput) []string
}

type MainStringListBuilder struct{}
