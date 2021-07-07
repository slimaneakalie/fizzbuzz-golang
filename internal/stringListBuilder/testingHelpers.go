package stringListBuilder

import (
	"encoding/json"
	"io/ioutil"
)

func LoadTestingData(path string) ([]StringListTestingDataElement, error) {
	jsonData, jsonReadingErr := ioutil.ReadFile(path)
	if jsonReadingErr != nil {
		return nil, jsonReadingErr
	}

	var testingData []StringListTestingDataElement

	unmarshallingJsonErr := json.Unmarshal(jsonData, &testingData)
	if unmarshallingJsonErr != nil {
		return nil, unmarshallingJsonErr
	}

	return testingData, nil
}

type StringListTestingDataElement struct {
	BuildInput     *BuildInput `json:"buildInput"`
	ExpectedOutput []string    `json:"expectedOutput"`
}
