package stringListBuilder

type StringListBuildInput struct {
	FirstInt  int
	SecondInt int
	Limit     int
	FirstStr  string
	SecondStr string
}

type defaultStringListBuilder struct{}
