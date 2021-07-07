package stringListBuilder

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("stringListBuilderTooling package - defaultStringListBuilder.go", func() {
	Context("defaultStringListBuilder.BuildStringList function", func() {
		It("should return the correct string list for each input", func() {
			testingData, err := loadTestingData()
			defaultBuilder := NewDefaultStringListBuilder()
			Expect(err).To(BeNil())

			for _, testElement := range testingData {
				actual := defaultBuilder.BuildStringList(testElement.BuildInput)
				// defaultBuilder.BuildStringList(testElement.BuildInput)
				Expect(actual).To(Equal(testElement.ExpectedOutput))
			}
		})

	})
})

func loadTestingData() ([]testingDataElement, error) {
	jsonData, jsonReadingErr := ioutil.ReadFile("../../test/assets/defaultStringListBuilderTesting.data.json")
	if jsonReadingErr != nil {
		return nil, jsonReadingErr
	}

	var testingData []testingDataElement

	unmarshallingJsonErr := json.Unmarshal(jsonData, &testingData)
	if unmarshallingJsonErr != nil {
		return nil, unmarshallingJsonErr
	}

	return testingData, nil
}

type testingDataElement struct {
	BuildInput     *StringListBuildInput `json:"buildInput"`
	ExpectedOutput []string              `json:"expectedOutput"`
}
