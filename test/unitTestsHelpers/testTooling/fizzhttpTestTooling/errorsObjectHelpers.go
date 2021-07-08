package fizzhttpTestTooling

import (
	"encoding/json"
	"reflect"
)

func NewRequestUnmarshalErrorType(targetedFieldName, requestWrongInputType string, expectedInputSample interface{}) *json.UnmarshalTypeError {
	return &json.UnmarshalTypeError{
		Field: targetedFieldName,
		Value: requestWrongInputType,
		Type:  reflect.TypeOf(expectedInputSample),
	}
}
