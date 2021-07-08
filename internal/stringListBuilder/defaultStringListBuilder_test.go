package stringListBuilder

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/common"
)

var _ = Describe("stringListBuilderTooling package - defaultStringListBuilder.go", func() {
	Context("defaultStringListBuilder.BuildStringList function", func() {
		It("should return the correct string list for each input", func() {
			var testingData []StringListTestingDataElement
			err := common.LoadTestingJsonData("../../test/assets/defaultStringListBuilderTesting.data.json", &testingData)
			Expect(err).To(BeNil())

			defaultBuilder := NewDefaultBuilder()
			for _, testElement := range testingData {
				actual := defaultBuilder.BuildStringList(testElement.BuildInput)
				expected := common.ToStringSlice(testElement.ExpectedOutput.([]interface{}))
				Expect(actual).To(Equal(expected))
			}
		})

	})
})
