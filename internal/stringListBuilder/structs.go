package stringListBuilder

type StringListBuildInput struct {
	FirstInt  int    `json:"firstInt"`
	SecondInt int    `json:"secondInt"`
	Limit     int    `json:"limit"`
	FirstStr  string `json:"firstStr"`
	SecondStr string `json:"secondStr"`
}

type defaultStringListBuilder struct{}
