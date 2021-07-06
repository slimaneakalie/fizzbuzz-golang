package fizzhttp

import (
	"github.com/go-playground/validator/v10"
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
			expected := "expected type string instead of int"
			unmarshalErrorExample := fizzhttpTestTooling.NewRequestUnmarshalErrorType("fieldName", "string", 123)
			actual := generateUnmarshalErrorResponseMessage(unmarshalErrorExample)
			Expect(actual).To(Equal(expected))
		})
	})

	Context("createValidationErrorServerResponse function", func() {
		It("should create the correct error http response based on the validation errors of the user query", func() {
			expected := NewHttpErrorResponseMetadata(BadRequestResponseTypeCode, nil)
			sampleErrors := validator.ValidationErrors{}

			actual := createValidationErrorServerResponse(sampleErrors)
			Expect(actual).To(Equal(expected))
		})
	})
})
