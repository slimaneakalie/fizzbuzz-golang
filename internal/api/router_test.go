package api

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/slimaneakalie/fizzbuzz-golang/internal/fizzbuzz"
	"github.com/slimaneakalie/fizzbuzz-golang/test/mocks/fizzhttpMocks"
)

var _ = Describe("Api package - Router.go", func() {
	Context("NewRouter function", func() {
		It("should create a new router while using the injected httpEngineFactory and stringListBuilder", func() {
			testingEngine := &fizzhttpMocks.TestingEngine{
				RouterGroups: []*fizzhttpMocks.TestingRouterGroup{},
			}

			testingEngineFactory := fizzhttpMocks.NewTestingEngineFactory(testingEngine)
			stringListBuilder := &fizzbuzz.MainStringListBuilder{}
			router := NewRouter(testingEngineFactory, stringListBuilder)

			Expect(testingEngineFactory.IsCalledNTimes("NewEngine", 1)).To(BeTrue())
			Expect(testingEngine.IsCalledWithParamsExactly("Group", "/v1/fizzbuzz")).To(BeTrue())
			Expect(testingEngine.RouterGroups).To(HaveLen(1))
			Expect(testingEngine.RouterGroups[0].IsCalledWithParamsPartially("POST", "/")).To(BeTrue())
			Expect(router.httpEngine).To(Equal(testingEngine))
		})
	})
})
