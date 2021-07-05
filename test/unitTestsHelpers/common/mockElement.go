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

func (mockElement *MockElement) GetNumberOfFuncCalls(funcName string) int {
	callData, _ := mockElement.funcCallsMap[funcName]
	return len(callData)
}

func (mockElement *MockElement) FuncIsCalledFirstTimeWithParamsPartial(funcName string, params ...interface{}) bool {
	funcFirstCallParams := mockElement.GetFuncFirstCallParamsInOrder(funcName)
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

func (mockElement *MockElement) GetFuncFirstCallParamsInOrder(funcName string) []interface{} {
	callData, ok := mockElement.funcCallsMap[funcName]
	if !ok || callData == nil {
		return nil
	}

	return callData[0].callParamsInOrder
}
