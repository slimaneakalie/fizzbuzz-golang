package stringListBuilder

import (
	"encoding/json"
	"fmt"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("stringListBuilderTooling package - defaultStringListBuilder.go", func() {
	Context("defaultStringListBuilder.BuildStringList function", func() {
		It("should return the correct string list for each input", func() {
			_, err := loadTestingData()
			Expect(err).To(BeNil())
			// Expect(testingData).To(Equal(12))
		})

	})
})

func loadTestingData() ([]testingDataElement, error) {
	jsonData, jsonReadingErr := ioutil.ReadFile("../../test/assets/defaultStringListBuilderTesting.data.json")
	if jsonReadingErr != nil {
		return nil, jsonReadingErr
	}

	testingData := &[]testingDataElement{}

	unmarshallingJsonErr := json.Unmarshal([]byte(jsonData), testingData)
	if unmarshallingJsonErr != nil {
		return nil, unmarshallingJsonErr
	}

	fmt.Println("testingData", *testingData)

	return *testingData, nil
}

type testingDataElement struct {
	buildInput     *StringListBuildInput `json:"buildInput" binding:"required"`
	expectedOutput []string              `json:"expectedOutput" binding:"required"`
}
