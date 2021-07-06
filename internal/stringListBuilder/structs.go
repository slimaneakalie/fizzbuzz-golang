package stringListBuilder

type StringListBuildInput struct {
	FirstInt  int    `json:"firstInt" binding:"required"`
	SecondInt int    `json:"secondInt" binding:"required"`
	Limit     int    `json:"limit" binding:"required"`
	FirstStr  string `json:"firstStr" binding:"required"`
	SecondStr string `json:"secondStr" binding:"required"`
}

type defaultStringListBuilder struct{}
