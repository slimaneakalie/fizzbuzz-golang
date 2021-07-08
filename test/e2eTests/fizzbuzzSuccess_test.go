package e2e

import (
	"encoding/json"
	"net/http"

	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/common"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/stringListBuilder"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/service"

	"github.com/slimaneakalie/fizzbuzz-golang/internal/api"
)

var _ = Describe("End to End tests - fizzbuzz success", func() {
	It("should send a valid string list when provided with a valid query", func() {
		var testingData []stringListBuilder.StringListTestingDataElement
		loadingErr := common.LoadTestingJsonData("../assets/defaultStringListBuilderTesting.data.json", &testingData)
		Expect(loadingErr).To(BeNil())
		expectedStatus := http.StatusOK

		for _, testingElement := range testingData {
			buildInput := testingElement.BuildInput
			query := api.FizzbuzzAPIRequest{
				FirstInt:  buildInput.FirstInt,
				SecondInt: buildInput.SecondInt,
				FirstStr:  buildInput.FirstStr,
				SecondStr: buildInput.SecondStr,
				Limit:     buildInput.Limit,
			}

			httpTestingEngine := service.New().HttpEngine
			responseRecorder, httpErr := performPOSTJSONRequest(httpTestingEngine, "/v1/fizzbuzz", &query)

			Expect(httpErr).To(BeNil())

			Expect(responseRecorder.Code).To(Equal(expectedStatus))

			var actualResponse api.FizzbuzzAPIResponse

			umarshallingErr := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)
			Expect(umarshallingErr).To(BeNil())
			expected := common.ToStringSlice(testingElement.ExpectedOutput.([]interface{}))
			Expect(actualResponse.FizzbuzzStringList).To(Equal(expected))
		}

	})

})
