package fizzhttp

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/testTooling/fizzhttpTestTooling"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("fizzhttp package - errors.go", func() {
	Context("createUnmarshalErrorServerResponse function", func() {
		It("should map the requestUnmarshalError correctly", func() {
			unmarshalErrorExample := fizzhttpTestTooling.NewRequestUnmarshalErrorType("fieldName", "string", 123)
			sampleErrorDetail := generateUnmarshalErrorResponseMessage(unmarshalErrorExample)
			sampleFieldsErrorr := []*responseFieldError{NewFieldError(ParamTypeMismatchErrorCode, unmarshalErrorExample.Field, sampleErrorDetail)}
			expected := NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, sampleFieldsErrorr)

			actual := createUnmarshalErrorServerResponse(unmarshalErrorExample)

			Expect(actual).To(Equal(expected))
		})
	})

	Context("generateUnmarshalErrorResponseMessage function", func() {
		It("should return the correct message based on the unmarshalling error", func() {
			expected := "expected type string instead of int"
			unmarshalErrorExample := fizzhttpTestTooling.NewRequestUnmarshalErrorType("fieldName", "string", 123)
			actual := generateUnmarshalErrorResponseMessage(unmarshalErrorExample)
			Expect(actual).To(Equal(expected))
		})
	})

	Context("createValidationErrorServerResponse function", func() {
		It("should create the correct error http response based on the validation errors of the user query", func() {
			fieldErrorsSample := []fieldValidationError{
				fizzhttpTestTooling.NewTestingFieldValidationError("required", "field1"),
				fizzhttpTestTooling.NewTestingFieldValidationError("", "field2"),
				fizzhttpTestTooling.NewTestingFieldValidationError("required", "field3"),
			}

			expectedResponseErrors := []*responseFieldError{
				NewFieldError(InvalidParamValueErrorCode, fieldErrorsSample[0].Field(), generateFieldErrorResponseMessage(fieldErrorsSample[0])),
				NewFieldError(InvalidParamValueErrorCode, fieldErrorsSample[1].Field(), generateFieldErrorResponseMessage(fieldErrorsSample[1])),
				NewFieldError(InvalidParamValueErrorCode, fieldErrorsSample[2].Field(), generateFieldErrorResponseMessage(fieldErrorsSample[2])),
			}
			expectedErrorServerResponse := NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, expectedResponseErrors)

			actualErrorServerResponse := createValidationErrorServerResponse(fieldErrorsSample)
			Expect(actualErrorServerResponse).To(Equal(expectedErrorServerResponse))
		})
	})
})
