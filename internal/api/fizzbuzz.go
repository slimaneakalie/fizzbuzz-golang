package api

import (
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzhttp"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"
)

func NewDefaultFizzbuzzRequestAPIHandler(stringListBuilder stringListBuilder.StringListBuilder, bindingErrorFormatter fizzhttp.BindingErrorFormatter) FizzbuzzRequestAPIHandler {
	return &defaultFizzbuzzRequestAPIHandler{
		StringListBuilder:     stringListBuilder,
		bindingErrorFormatter: bindingErrorFormatter,
	}
}

func (handler *defaultFizzbuzzRequestAPIHandler) handleFizzbuzzRequest() fizzhttp.HandlerFunc {
	return handler.fizzbuzzRequestHandler
}

func (handler *defaultFizzbuzzRequestAPIHandler) fizzbuzzRequestHandler(context fizzhttp.RequestContext) {
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

func toFizzbuzzListBuildInput(request *FizzbuzzAPIRequest) *stringListBuilder.BuildInput {
	return &stringListBuilder.BuildInput{
		FirstInt:  request.FirstInt,
		SecondInt: request.SecondInt,
		Limit:     request.Limit,
		FirstStr:  request.FirstStr,
		SecondStr: request.SecondStr,
	}
}
