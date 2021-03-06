package api

import (
	"errors"

	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/testTooling/apiTestTooling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api package - Router.go", func() {
	Context("Router.NewRouter function", func() {
		It("should create a new Router while using the injected httpEngineFactory and stringListBuilderTooling", func() {
			testTooling := apiTestTooling.PrepareRouterTestsTooling(nil)

			router := NewRouter(testTooling.TestingEngineFactory, testTooling.TestingStringListBuilder, testTooling.TestingMonitoringHandler, testTooling.TestingLogger)

			Expect(testTooling.TestingEngineFactory.GetNumberOfFuncCalls("NewEngine")).To(Equal(1))
			Expect(testTooling.TestingEngine.GetNumberOfFuncCalls("GET")).To(Equal(1))
			Expect(testTooling.TestingEngine.GetNumberOfFuncCalls("UseMiddleware")).To(Equal(1))
			Expect(testTooling.TestingMonitoringHandler.GetNumberOfFuncCalls("HandleMonitoringQuery")).To(Equal(1))

			expectedGroupFuncCallParam := []interface{}{"/v1/fizzbuzz"}
			Expect(testTooling.TestingEngine.GetFuncFirstCallParamsInOrder("Group")).To(Equal(expectedGroupFuncCallParam))
			Expect(testTooling.TestingEngine.RouterGroups).To(HaveLen(1))
			Expect(testTooling.TestingEngine.RouterGroups[0].FuncIsCalledFirstTimeWithParamsPartial("POST", "")).To(BeTrue())
			Expect(testTooling.TestingEngine.RouterGroups[0].FuncIsCalledFirstTimeWithParamsPartial("GET", "/stats")).To(BeTrue())

			Expect(router.HttpEngine).To(Equal(testTooling.TestingEngine))
			Expect(router.logger).To(Equal(testTooling.TestingLogger))
		})
	})

	Context("Router.Run function", func() {
		It("should start the http server and log the port", func() {
			testTooling := apiTestTooling.PrepareRouterTestsTooling(nil)
			testRunRouterCommonExecutionPath(testTooling)
			Expect(testTooling.TestingLogger.GetNumberOfFuncCalls("Error")).To(Equal(0))
		})

		It("should start the http server, log the port and log the error returned by the http engine", func() {
			httpEngineRunMethodOutput := errors.New("port is already in use")
			testTooling := apiTestTooling.PrepareRouterTestsTooling(httpEngineRunMethodOutput)
			testRunRouterCommonExecutionPath(testTooling)

			expectedErrorFuncCallParams := []interface{}{"Error running server caused by:", httpEngineRunMethodOutput.Error()}
			Expect(testTooling.TestingLogger.GetFuncFirstCallParamsInOrder("Error")).To(Equal(expectedErrorFuncCallParams))
			Expect(testTooling.TestingLogger.GetNumberOfFuncCalls("Error")).To(Equal(1))
		})

	})
})

func testRunRouterCommonExecutionPath(testTooling *apiTestTooling.Router) {
	router := NewRouter(testTooling.TestingEngineFactory, testTooling.TestingStringListBuilder, testTooling.TestingMonitoringHandler, testTooling.TestingLogger)
	port := 9000
	router.Run(port)
	expectedInfoFuncCallParams := []interface{}{"Running server on port", port}
	Expect(testTooling.TestingLogger.GetFuncFirstCallParamsInOrder("Info")).To(Equal(expectedInfoFuncCallParams))
	Expect(testTooling.TestingLogger.GetNumberOfFuncCalls("Info")).To(Equal(1))

	expectedRunFuncCallParams := []interface{}{port}
	Expect(testTooling.TestingEngine.GetFuncFirstCallParamsInOrder("Run")).To(Equal(expectedRunFuncCallParams))
	Expect(testTooling.TestingEngine.GetNumberOfFuncCalls("Run")).To(Equal(1))

}
