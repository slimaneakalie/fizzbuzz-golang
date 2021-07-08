package stringListBuilder

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("stringListBuilderTooling package - defaultStringListBuilder.go", func() {
	Context("defaultStringListBuilder.BuildStringList function", func() {
		It("should return the correct string list for each input", func() {
			testingData, err := LoadTestingData("../../test/assets/defaultStringListBuilderTesting.data.json")
			Expect(err).To(BeNil())

			defaultBuilder := NewDefaultBuilder()
			for _, testElement := range testingData {
				actual := defaultBuilder.BuildStringList(testElement.BuildInput)
				// defaultBuilder.BuildStringList(testElement.BuildInput)
				Expect(actual).To(Equal(testElement.ExpectedOutput))
			}
		})

	})
})
