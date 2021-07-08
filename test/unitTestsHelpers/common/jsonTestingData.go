package common

import (
	"encoding/json"
	"io/ioutil"
)

func LoadTestingJsonData(path string, targetObject interface{}) error {
	jsonData, jsonReadingErr := ioutil.ReadFile(path)
	if jsonReadingErr != nil {
		return jsonReadingErr
	}

	return json.Unmarshal(jsonData, targetObject)
}

func ToStringSlice(inputSlice []interface{}) []string {
	var outputSlice []string
	for _, inputElement := range inputSlice {
		outputSlice = append(outputSlice, inputElement.(string))
	}

	return outputSlice
}
