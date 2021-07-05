package logger

import "fmt"

func NewDefaultLogger() Logger {
	return &defaultLogger{}
}

func (logger *defaultLogger) Info(data ...interface{}) {
	fmt.Println(data...)
}

func (logger *defaultLogger) Error(data ...interface{}) {
	fmt.Println(data...)
}
