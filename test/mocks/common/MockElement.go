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

func (mockElement *MockElement) IsCalledWith(funcName string, params []interface{}) bool {
	callData, ok := mockElement.funcCallsMap[funcName]
	if !ok || len(callData) == 0 {
		return false
	}

	for index, element := range callData[0].callParamsInOrder {
		if params[index] != element {
			return false
		}
	}

	return true
}
