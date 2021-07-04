package api

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/testTooling"
)

var _ = Describe("Api package - Router.go", func() {
	Context("NewRouter function", func() {
		It("should create a new router while using the injected httpEngineFactory and stringListBuilder", func() {
			tooling := testTooling.PrepareRouterTestsTooling(nil)

			router := NewRouter(tooling.TestingEngineFactory, tooling.StringListBuilder)

			Expect(tooling.TestingEngineFactory.IsCalledNTimes("NewEngine", 1)).To(BeTrue())
			Expect(tooling.TestingEngine.IsCalledWithParamsExactly("Group", "/v1/fizzbuzz")).To(BeTrue())
			Expect(tooling.TestingEngine.RouterGroups).To(HaveLen(1))
			Expect(tooling.TestingEngine.RouterGroups[0].IsCalledWithParamsPartially("POST", "/")).To(BeTrue())
			Expect(router.httpEngine).To(Equal(tooling.TestingEngine))
		})
	})

	Context("Router.Run function", func() {
		It("should start server and print the port", func() {
			tooling := testTooling.PrepareRouterTestsTooling(nil)

			router := NewRouter(tooling.TestingEngineFactory, tooling.StringListBuilder)
			port := 9000
			router.Run(port)

			Expect(tooling.TestingEngine.IsCalledNTimes("Run", 1)).To(BeTrue())
			Expect(tooling.TestingEngine.IsCalledWithParamsExactly("Run", port)).To(BeTrue())
		})

	})
})
