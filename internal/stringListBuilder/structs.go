package stringListBuilder

type BuildInput struct {
	FirstInt  int    `json:"firstInt"`
	SecondInt int    `json:"secondInt"`
	Limit     int    `json:"limit"`
	FirstStr  string `json:"firstStr"`
	SecondStr string `json:"secondStr"`
}

type defaultStringListBuilder struct{}

type StringListTestingDataElement struct {
	BuildInput     *BuildInput `json:"buildInput"`
	ExpectedOutput interface{} `json:"expectedOutput"`
}
