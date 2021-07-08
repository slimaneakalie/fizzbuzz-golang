package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
)

func performPOSTJSONRequest(handler http.Handler, route string, queryObject *api.FizzbuzzAPIRequest) (*httptest.ResponseRecorder, error) {
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
