package e2e

import (
	"net/http"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("End to End tests - fizzbuzz bad requests", func() {
	It("should send the appropriate response for each bad request", func() {
		testInput := &multipleRequestsTestInput{
			testJsonDataPath:           "../assets/e2eBadRequestsTesting.data.json",
			expectedHttpResponseStatus: http.StatusBadRequest,
		}

		testMultipleE2ERequests(testInput)
	})

})
