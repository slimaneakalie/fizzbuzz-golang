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
			httpEngineFactory := fizzhttpMocks.NewTestEngineFactory()
			stringListBuilder := &fizzbuzz.MainStringListBuilder{}
			router := NewRouter(httpEngineFactory, stringListBuilder)
			// httpEngineForTest := MockElement
			// httpRouterGroupForTest := MockElement
			// Expect(httpEngineFactory.isCalled('NewEngine')).To(BeTrue())
			Expect(httpEngineFactory.IsCalledNTimes("NewEngine", 1)).To(BeTrue())
			Expect(router.httpEngine).To(Equal(httpEngineForTest))
			Expect(method.isCalledWith(httpEngineForTest.Group, "/v1/fizzbuzz")).To(BeTrue())
			Expect(httpRouterGroupForTest.POSTEndpoints).To(Equal([]string{"/"}))
		})
	})
})
