package common

func (mockElement *MockElement) recordFuncCall(funcName string, callParamsInOrder ...interface{}) {
	callData := FuncCallData{
		callParamsInOrder: callParamsInOrder,
	}

	if currentCalls, ok := mockElement.funcCallsMap[funcName]; ok {
		mockElement.funcCallsMap[funcName] = append(currentCalls, callData)
	} else {
		mockElement.funcCallsMap[funcName] = []FuncCallData{callData}
	}
}
