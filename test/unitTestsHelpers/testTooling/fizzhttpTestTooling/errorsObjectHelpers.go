package fizzhttpTestTooling

import (
	"encoding/json"
	"reflect"
)

func NewRequestUnmarshalErrorType(targetedFieldName, jsonTypeDescription string, requestWrongInput interface{}) *json.UnmarshalTypeError {
	return &json.UnmarshalTypeError{
		Field: targetedFieldName,
		Value: jsonTypeDescription,
		Type:  reflect.TypeOf(requestWrongInput),
	}
}
