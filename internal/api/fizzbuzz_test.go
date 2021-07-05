package api

import (
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/testTooling/apiTestTooling"
)

var _ = Describe("Api package - fizzbuzz.go", func() {
	Context("handler.fizzbuzzRequestHandler function", func() {
		It("should send JSON successful response if there is no error in user request parsing", func() {
			finalStringList := []string{"this", "is", "a", "stub"}
			testTooling := apiTestTooling.PrepareFizzbuzzTestsTooling(finalStringList, nil, nil)

			invokeFizzbuzzRequestHandler(testTooling)

			expectedApiResponse := fizzbuzzAPIResponse{
				FizzbuzzStringList: finalStringList,
			}
			expectedJSONFuncCallParams := []interface{}{http.StatusOK, expectedApiResponse}
			Expect(testTooling.TestingRequestContext.GetFuncFirstCallParamsInOrder("JSON")).To(Equal(expectedJSONFuncCallParams))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("JSON")).To(Equal(1))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("AbortWithStatusJSON")).To(Equal(0))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("ShouldBindBodyWith")).To(Equal(1))

			var userRequest fizzbuzzAPIRequest
			expectedBuildStringListCallParams := []interface{}{toFizzbuzzListBuildInput(&userRequest)}
			Expect(testTooling.TestingStringListBuilder.GetFuncFirstCallParamsInOrder("BuildStringList")).To(Equal(expectedBuildStringListCallParams))
			Expect(testTooling.TestingStringListBuilder.GetNumberOfFuncCalls("BuildStringList")).To(Equal(1))
		})

		It("should abort with bad request status if there is an error in user request parsing", func() {
			shouldBindBodyWithOutput := errors.New("JSON parsing error example")
			formatBindingErrorOutput := errors.New("binding error format example")
			testTooling := apiTestTooling.PrepareFizzbuzzTestsTooling(nil, shouldBindBodyWithOutput, formatBindingErrorOutput)
			invokeFizzbuzzRequestHandler(testTooling)

			expectedAbortFuncCallParams := []interface{}{http.StatusBadRequest, formatBindingErrorOutput}
			Expect(testTooling.TestingRequestContext.GetFuncFirstCallParamsInOrder("AbortWithStatusJSON")).To(Equal(expectedAbortFuncCallParams))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("AbortWithStatusJSON")).To(Equal(1))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("JSON")).To(Equal(0))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("ShouldBindBodyWith")).To(Equal(1))

			expectedFormatBindingErrorCallParams := []interface{}{shouldBindBodyWithOutput}
			Expect(testTooling.TestingEngine.GetFuncFirstCallParamsInOrder("FormatBindingError")).To(Equal(expectedFormatBindingErrorCallParams))
			Expect(testTooling.TestingEngine.GetNumberOfFuncCalls("FormatBindingError")).To(Equal(1))

			Expect(testTooling.TestingStringListBuilder.GetNumberOfFuncCalls("BuildStringList")).To(Equal(0))
		})
	})
})

func invokeFizzbuzzRequestHandler(testTooling *apiTestTooling.Fizzbuzz) {
	apiHandler := NewDefaultFizzbuzzRequestAPIHandler(testTooling.TestingStringListBuilder, testTooling.TestingEngine.FormatBindingError)
	requestHandler := apiHandler.handleFizzbuzzRequest()
	requestHandler(testTooling.TestingRequestContext)
}
