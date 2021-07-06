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
			sampleFieldsErrorr := []*fieldError{NewFieldError(ParamTypeMismatchErrorCode, unmarshalErrorExample.Field, sampleErrorDetail)}
			expected := NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, sampleFieldsErrorr)

			actual := createUnmarshalErrorServerResponse(unmarshalErrorExample)

			Expect(actual).To(Equal(expected))
		})
	})

	Context("generateUnmarshalErrorResponseMessage function", func() {
		It("should return the correct message based on the unmarshalling error", func() {
			unmarshalErrorExample := fizzhttpTestTooling.NewRequestUnmarshalErrorType("fieldName", "string", 123)
			expected := "expected type string instead of int"
			actual := generateUnmarshalErrorResponseMessage(unmarshalErrorExample)
			Expect(actual).To(Equal(expected))
		})
	})
})
