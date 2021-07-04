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
			testingEngine := &fizzhttpMocks.TestingEngine{}
			testingEngineFactory := fizzhttpMocks.NewTestingEngineFactory(testingEngine)
			stringListBuilder := &fizzbuzz.MainStringListBuilder{}
			router := NewRouter(testingEngineFactory, stringListBuilder)
			// httpRouterGroupForTest := MockElement
			Expect(testingEngineFactory.IsCalledNTimes("NewEngine", 1)).To(BeTrue())
			Expect(router.httpEngine).To(Equal(testingEngine))
			Expect(testingEngine.IsCalledWith("Group", []interface{}{"/v1/fizzbuzz"})).To(BeTrue())
			// Expect(httpRouterGroupForTest.POSTEndpoints).To(Equal([]string{"/"}))
		})
	})
})
