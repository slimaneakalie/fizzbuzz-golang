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
