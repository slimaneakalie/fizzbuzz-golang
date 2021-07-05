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
		//It("should send JSON successful response if there is no error in user request parsing", func() {
		//	testTooling := apiTestTooling.PrepareFizzbuzzTestsTooling(nil, nil)
		//	invokeFizzbuzzRequestHandler(testTooling)
		//})

		It("should abort with bad request status if there is an error in user request parsing", func() {
			shouldBindBodyWithOutput := errors.New("JSON parsing error example")
			formatBindingErrorOutput := errors.New("binding error format example")
			testTooling := apiTestTooling.PrepareFizzbuzzTestsTooling(shouldBindBodyWithOutput, formatBindingErrorOutput)
			invokeFizzbuzzRequestHandler(testTooling)

			expectedAbortFuncCallParams := []interface{}{http.StatusBadRequest, formatBindingErrorOutput}
			Expect(testTooling.TestingRequestContext.GetFuncFirstCallParamsInOrder("AbortWithStatusJSON")).To(Equal(expectedAbortFuncCallParams))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("AbortWithStatusJSON")).To(Equal(1))
			Expect(testTooling.TestingRequestContext.GetNumberOfFuncCalls("JSON")).To(Equal(0))

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

//func createTestTooling() (*apiTestTooling.Fizzbuzz, *defaultFizzbuzzRequestAPIHandler) {
//	testTooling := apiTestTooling.PrepareFizzbuzzTestsTooling()
//	handler := NewDefaultFizzbuzzRequestAPIHandler(testTooling.TestingStringListBuilder, testTooling.TestingEngine.FormatBindingError).(*defaultFizzbuzzRequestAPIHandler)
//	return testTooling, handler
//}
