package e2e

type multipleRequestsTestInput struct {
	testJsonDataPath           string
	expectedHttpResponseStatus int
}

type testingData struct {
	Request          interface{} `json:"request"`
	ExpectedResponse interface{} `json:"expectedResponse"`
}
