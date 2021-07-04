package common

type FuncCallData struct {
	callParamsInOrder []interface{}
}

type FuncCallsInOrder []FuncCallData

type MockElement struct {
	funcCallsMap map[string]FuncCallsInOrder
}
