package logger

type Logger interface {
	Info(data ...interface{})
	Error(data ...interface{})
}
