package logger

import "fmt"

func NewMainLogger() Logger {
	return &mainLogger{}
}

func (logger *mainLogger) Info(data ...interface{}) {
	fmt.Println(data...)
}

func (logger *mainLogger) Error(data ...interface{}) {
	fmt.Println(data...)
}
