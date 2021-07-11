package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/monitoringMocks"

	. "github.com/onsi/gomega"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/service"
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/common"
)

func performPOSTJSONRequest(handler http.Handler, route string, queryObject interface{}) (*httptest.ResponseRecorder, error) {
	var req *http.Request

	unmarshalledObject, err := json.Marshal(queryObject)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(unmarshalledObject)
	req, err = http.NewRequest("POST", route, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	responseRecorder := httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder, nil
}

func testMultipleE2ERequests(testInput *multipleRequestsTestInput) {
	var testData []testingData
	loadingErr := common.LoadTestingJsonData(testInput.testJsonDataPath, &testData)
	Expect(loadingErr).To(BeNil())

	testingMonitoringHandler := monitoringMocks.NewTestingHandler()
	httpTestingEngine := service.NewServer(testingMonitoringHandler).HttpEngine

	for _, testingElement := range testData {
		responseRecorder, httpErr := performPOSTJSONRequest(httpTestingEngine, "/v1/fizzbuzz", &testingElement.Request)

		Expect(httpErr).To(BeNil())

		Expect(responseRecorder.Code).To(Equal(testInput.expectedHttpResponseStatus))

		var actual interface{}

		unmarshallingErr := json.Unmarshal(responseRecorder.Body.Bytes(), &actual)
		Expect(unmarshallingErr).To(BeNil())

		Expect(actual).To(Equal(testingElement.ExpectedResponse))
	}
}
