package api

import (
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzBuzz"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (handler *MainFizzBuzzRequestAPIHandler) HandleFizzBuzzRequest() gin.HandlerFunc {
	return handler.fizzBuzzRequestHandler
}

func (handler *MainFizzBuzzRequestAPIHandler) fizzBuzzRequestHandler(c *gin.Context) {
	var userRequest FizzBuzzAPIRequest
	if err := c.ShouldBindBodyWith(&userRequest, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, handleBindError(err))
		return
	}

	fizzBuzzListBuildInput := toFizzBuzzListBuildInput(&userRequest)
	fizzBuzzStringList := handler.StringListBuilder.BuildStringList(fizzBuzzListBuildInput)

	apiResponse := FizzBuzzAPIResponse{
		fizzBuzzStringList: fizzBuzzStringList,
	}

	c.JSON(http.StatusOK, apiResponse)
}

func toFizzBuzzListBuildInput(request *FizzBuzzAPIRequest) *fizzBuzz.StringListBuildInput {
	return &fizzBuzz.StringListBuildInput{
		FirstInt:  request.firstInt,
		SecondInt: request.secondInt,
		ThirdInt:  request.thirdInt,
		Limit:     request.limit,
		FirstStr:  request.firstStr,
		SecondStr: request.secondStr,
	}
}
