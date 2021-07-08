package e2e

import (
	"net/http"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("End to End tests - fizzbuzz success", func() {
	It("should send a valid string list when provided with a valid query", func() {
		testInput := &multipleRequestsTestInput{
			testJsonDataPath:           "../assets/e2eSuccessRequestsTesting.data.json",
			expectedHttpResponseStatus: http.StatusOK,
		}

		testMultipleE2ERequests(testInput)
	})
})
