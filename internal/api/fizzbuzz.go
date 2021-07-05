package api

import (
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

func NewMainFizzbuzzRequestAPIHandler(stringListBuilder stringListBuilder.StringListBuilder, bindingErrorFormatter fizzhttp.BindingErrorFormatter) FizzbuzzRequestAPIHandler {
	return &mainFizzbuzzRequestAPIHandler{
		StringListBuilder:     stringListBuilder,
		bindingErrorFormatter: bindingErrorFormatter,
	}
}

func (handler *mainFizzbuzzRequestAPIHandler) handleFizzbuzzRequest() fizzhttp.HandlerFunc {
	return handler.fizzbuzzRequestHandler
}

func (handler *mainFizzbuzzRequestAPIHandler) fizzbuzzRequestHandler(context fizzhttp.RequestContext) {
	var userRequest fizzbuzzAPIRequest
	if err := context.ShouldBindBodyWith(&userRequest, fizzhttp.JsonBinding); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, handler.bindingErrorFormatter(err))
		return
	}

	fizzbuzzListBuildInput := toFizzbuzzListBuildInput(&userRequest)
	fizzbuzzStringList := handler.StringListBuilder.BuildStringList(fizzbuzzListBuildInput)

	apiResponse := fizzbuzzAPIResponse{
		FizzbuzzStringList: fizzbuzzStringList,
	}

	context.JSON(http.StatusOK, apiResponse)
}

func toFizzbuzzListBuildInput(request *fizzbuzzAPIRequest) *stringListBuilder.StringListBuildInput {
	return &stringListBuilder.StringListBuildInput{
		FirstInt:  request.FirstInt,
		SecondInt: request.SecondInt,
		Limit:     request.Limit,
		FirstStr:  request.FirstStr,
		SecondStr: request.SecondStr,
	}
}
