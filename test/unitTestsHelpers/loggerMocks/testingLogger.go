package loggerMocks

func NewTestingLogger() *TestingLogger {
	return &TestingLogger{}
}

func (logger *TestingLogger) Info(data ...interface{}) {
	logger.RecordFuncCall("Info", data...)
}

func (logger *TestingLogger) Error(data ...interface{}) {
	logger.RecordFuncCall("Error", data...)
}
