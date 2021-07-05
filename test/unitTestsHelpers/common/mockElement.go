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

func (mockElement *MockElement) FuncIsCalledExactlyNTimes(funcName string, n int) bool {
	callData, _ := mockElement.funcCallsMap[funcName]
	return callData != nil && len(callData) == n
}

func (mockElement *MockElement) FuncIsCalledFirstTimeWithParamsExact(funcName string, params ...interface{}) bool {
	funcFirstCallParams := mockElement.getFuncFirstCallParamsInOrder(funcName)
	for index, element := range funcFirstCallParams {
		if params[index] != element {
			return false
		}
	}

	return true
}

func (mockElement *MockElement) FuncIsCalledFirstTimeWithParamsPartial(funcName string, params ...interface{}) bool {
	funcFirstCallParams := mockElement.getFuncFirstCallParamsInOrder(funcName)
	if funcFirstCallParams == nil {
		return false
	}

	for index, element := range params {
		if funcFirstCallParams[index] != element {
			return false
		}
	}

	return true
}

func (mockElement *MockElement) getFuncFirstCallParamsInOrder(funcName string) []interface{} {
	callData, ok := mockElement.funcCallsMap[funcName]
	if !ok || callData == nil {
		return nil
	}

	return callData[0].callParamsInOrder
}
