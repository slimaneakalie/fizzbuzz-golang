package fizzhttp

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("fizzhttp package - errorObjectHelpers.go", func() {
	Context("NewHttpErrorResponseMetadata function", func() {
		It("should map params to a valid httpErrorResponseMetadata object", func() {
			sampleFieldsError := []*responseFieldError{
				NewFieldError(ParamTypeMismatchErrorCode, "fieldName", "error detail"),
			}
			sampleErrorType := BadRequestResponseTypeCode

			expected := &httpErrorResponseMetadata{
				Type:        sampleErrorType,
				FieldErrors: sampleFieldsError,
			}
			actual := NewHttpErrorResponseMetadata(sampleErrorType, sampleFieldsError)

			Expect(actual).To(Equal(expected))
		})
	})
})
