package api

import (
	"errors"

	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/testTooling/apiTestTooling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api package - router.go", func() {
	Context("router.NewRouter function", func() {
		It("should create a new router while using the injected httpEngineFactory and stringListBuilder", func() {
			tooling := apiTestTooling.PrepareRouterTestsTooling(nil)

			router := NewRouter(tooling.TestingEngineFactory, tooling.StringListBuilder, tooling.TestingLogger)

			Expect(tooling.TestingEngineFactory.FuncIsCalledExactlyNTimes("NewEngine", 1)).To(BeTrue())
			Expect(tooling.TestingEngine.FuncIsCalledFirstTimeWithParamsExact("Group", "/v1/stringListBuilder")).To(BeTrue())
			Expect(tooling.TestingEngine.RouterGroups).To(HaveLen(1))
			Expect(tooling.TestingEngine.RouterGroups[0].FuncIsCalledFirstTimeWithParamsPartial("POST", "/")).To(BeTrue())
			Expect(router.httpEngine).To(Equal(tooling.TestingEngine))
			Expect(router.logger).To(Equal(tooling.TestingLogger))
		})
	})

	Context("router.Run function", func() {
		It("should start the http server and log the port", func() {
			tooling := apiTestTooling.PrepareRouterTestsTooling(nil)
			testRunRouterCommonExecutionPath(tooling)
			Expect(tooling.TestingLogger.FuncIsCalledExactlyNTimes("Error", 0))
		})

		It("should start the http server, log the port and log the error returned by the http engine", func() {
			httpEngineRunMethodOutput := errors.New("port is already in use")
			tooling := apiTestTooling.PrepareRouterTestsTooling(httpEngineRunMethodOutput)
			testRunRouterCommonExecutionPath(tooling)

			Expect(tooling.TestingLogger.FuncIsCalledExactlyNTimes("Error", 1))
			Expect(tooling.TestingLogger.FuncIsCalledFirstTimeWithParamsExact("Error", "Error running server caused by:", httpEngineRunMethodOutput.Error()))
		})

	})
})

func testRunRouterCommonExecutionPath(tooling *apiTestTooling.RouterTestsTooling) {
	router := NewRouter(tooling.TestingEngineFactory, tooling.StringListBuilder, tooling.TestingLogger)
	port := 9000
	router.Run(port)

	Expect(tooling.TestingLogger.FuncIsCalledExactlyNTimes("Info", 1))
	Expect(tooling.TestingLogger.FuncIsCalledFirstTimeWithParamsExact("Info", "Running server on port", port))

	Expect(tooling.TestingEngine.FuncIsCalledExactlyNTimes("Run", 1)).To(BeTrue())
	Expect(tooling.TestingEngine.FuncIsCalledFirstTimeWithParamsExact("Run", port)).To(BeTrue())
}
