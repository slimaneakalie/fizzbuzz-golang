package fizzhttp

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/testTooling/fizzhttpTestTooling"

	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("fizzhttp package - defaultEngine.go", func() {
	Context("engine.FormatBindingError function", func() {
		It("should return the unmarshalling error format when provided with a requestUnmarshalErrorType error", func() {
			unmarshalErrorExample := &json.UnmarshalTypeError{
				Field: "field_name_example",
				Value: "int",
				Type:  reflect.TypeOf("a"),
			}

			expectedFormattedError := &gin.Error{
				Err:  unmarshalErrorExample,
				Type: gin.ErrorTypePublic,
				Meta: createUnmarshalErrorServerResponse(unmarshalErrorExample),
			}

			testBindingBindingErrorFormatting(unmarshalErrorExample, expectedFormattedError)
		})

		It("should return the validation errors format when provided with a validationErrorsType error", func() {
			validationErrorsExample := validationErrorsType{
				fizzhttpTestTooling.NewTestingFieldValidationError("required", "field1"),
			}

			expectedFormattedError := &gin.Error{
				Err:  validationErrorsExample,
				Type: gin.ErrorTypePublic,
				Meta: createValidationErrorServerResponse(validationErrorsExample),
			}

			testBindingBindingErrorFormatting(validationErrorsExample, expectedFormattedError)
		})

		It("should return the default format when provided with a generic error", func() {
			genericBindingErrorExample := errors.New("http error example")
			expectedFormattedError := &gin.Error{
				Err:  genericBindingErrorExample,
				Type: gin.ErrorTypePublic,
				Meta: httpErrorResponseMetadata{
					Type: BadRequestResponseTypeCode,
				},
			}

			testBindingBindingErrorFormatting(genericBindingErrorExample, expectedFormattedError)
		})
	})
})

func testBindingBindingErrorFormatting(bindingErrorExample, expectedFormattedError error) {
	httpEngineFactory := NewDefaultHttpEngineFactory(ReleaseMode)
	httpEngine := httpEngineFactory.NewEngine()

	actual := httpEngine.FormatBindingError(bindingErrorExample)
	Expect(actual).To(Equal(expectedFormattedError))
}
