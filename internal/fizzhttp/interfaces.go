package fizzhttp

import "net/http"

type EngineFactory interface {
	NewEngine() Engine
}

type Engine interface {
	RouterGroup
	Group(relativePath string) RouterGroup
	FormatBindingError(bindError error) error
	Run(port int) error
	ServeHTTP(http.ResponseWriter, *http.Request)
	UseMiddleware(handlerFunc HandlerFunc)
}

type RouterGroup interface {
	POST(relativePath string, handler HandlerFunc)
	GET(relativePath string, handler HandlerFunc)
}

type RequestContext interface {
	ShouldBindBodyWithJSON(targetObjectPointer interface{}) (err error)
	AbortWithStatusJSON(statusCode int, responseObject interface{})
	SendJSONResponse(code int, obj interface{})
	GetResponseWriter() http.ResponseWriter
	GetRequest() *http.Request
	GetResponseStatus() int
	GetJsonStringQuery() string
	Next()
}

type fieldValidationError interface {
	Tag() string
	Field() string
}
