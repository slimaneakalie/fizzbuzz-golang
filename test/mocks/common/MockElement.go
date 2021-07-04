package common

func (mockElement *MockElement) RecordFuncCall(funcName string, callParamsInOrder ...interface{}) {
	if mockElement.funcCallsMap == nil {
		mockElement.funcCallsMap = map[string]FuncCallsInOrder{}
	}

	callData := FuncCallData{
		callParamsInOrder: callParamsInOrder,
	}

	if currentCalls, ok := mockElement.funcCallsMap[funcName]; ok {
		mockElement.funcCallsMap[funcName] = append(currentCalls, callData)
	} else {
		mockElement.funcCallsMap[funcName] = []FuncCallData{callData}
	}
}

func (mockElement *MockElement) IsCalledNTimes(funcName string, n int) bool {
	callData, _ := mockElement.funcCallsMap[funcName]
	return callData != nil && len(callData) == n
}
