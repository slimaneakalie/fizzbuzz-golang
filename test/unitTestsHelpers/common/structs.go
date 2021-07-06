package common

type FuncCallData struct {
	callParamsInOrder []interface{}
}

type FuncCallsInOrder []FuncCallData

type MockElement struct {
	funcCallsMap map[string]FuncCallsInOrder
}

type TestingFieldValidationError struct {
	TagOutput   string
	FieldOutput string
}
