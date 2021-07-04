package api

import (
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzbuzz"
)

func (handler *MainFizzbuzzRequestAPIHandler) handleFizzbuzzRequest() fizzhttp.HandlerFunc {
	return handler.fizzbuzzRequestHandler
}

func (handler *MainFizzbuzzRequestAPIHandler) fizzbuzzRequestHandler(context fizzhttp.RequestContext) {
	var userRequest FizzbuzzAPIRequest
	if err := context.ShouldBindBodyWith(&userRequest, fizzhttp.JsonBinding); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, handler.bindingErrorFormatter(err))
		return
	}

	fizzbuzzListBuildInput := toFizzbuzzListBuildInput(&userRequest)
	fizzbuzzStringList := handler.StringListBuilder.BuildStringList(fizzbuzzListBuildInput)

	apiResponse := FizzbuzzAPIResponse{
		FizzbuzzStringList: fizzbuzzStringList,
	}

	context.JSON(http.StatusOK, apiResponse)
}

func toFizzbuzzListBuildInput(request *FizzbuzzAPIRequest) *fizzbuzz.StringListBuildInput {
	return &fizzbuzz.StringListBuildInput{
		FirstInt:  request.FirstInt,
		SecondInt: request.SecondInt,
		Limit:     request.Limit,
		FirstStr:  request.FirstStr,
		SecondStr: request.SecondStr,
	}
}
