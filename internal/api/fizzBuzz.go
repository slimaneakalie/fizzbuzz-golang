package api

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzBuzz"

	"github.com/gin-gonic/gin"
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
		FizzBuzzStringList: fizzBuzzStringList,
	}

	c.JSON(http.StatusOK, apiResponse)
}

func toFizzBuzzListBuildInput(request *FizzBuzzAPIRequest) *fizzBuzz.StringListBuildInput {
	return &fizzBuzz.StringListBuildInput{
		FirstInt:  request.FirstInt,
		SecondInt: request.SecondInt,
		Limit:     request.Limit,
		FirstStr:  request.FirstStr,
		SecondStr: request.SecondStr,
	}
}
