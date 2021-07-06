package fizzhttpTestTooling

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/common"
)

func NewTestingFieldValidationError(tag, field string) *common.TestingFieldValidationError {
	return &common.TestingFieldValidationError{
		TagOutput:   tag,
		FieldOutput: field,
	}
}
