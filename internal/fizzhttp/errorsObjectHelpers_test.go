package fizzhttp

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("fizzhttp package - errorObjectHelpers.go", func() {
	Context("NewHttpErrorResponse function", func() {
		It("should map params to a valid httpErrorResponse object", func() {
			sampleFieldsError := []*responseFieldError{
				NewFieldError(ParamTypeMismatchErrorCode, "fieldName", "error detail"),
			}
			sampleErrorType := BadRequestResponseTypeCode

			expected := &httpErrorResponse{
				Type:        sampleErrorType,
				FieldErrors: sampleFieldsError,
			}
			actual := NewHttpErrorResponse(sampleErrorType, sampleFieldsError)

			Expect(actual).To(Equal(expected))
		})
	})

	Context("NewFieldError function", func() {
		It("should map params to a valid responseFieldError object", func() {
			errorType, field, errorDetail := BadRequestResponseTypeCode, "limit", "invalid value"

			expected := &responseFieldError{
				Type:   errorType,
				Field:  field,
				Detail: errorDetail,
			}

			actual := NewFieldError(errorType, field, errorDetail)

			Expect(actual).To(Equal(expected))
		})
	})
})
