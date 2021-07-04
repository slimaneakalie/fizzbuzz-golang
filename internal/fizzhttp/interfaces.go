package fizzhttp

import "github.com/gin-gonic/gin/binding"

type EngineFactory interface {
	NewEngine() Engine
}

type Engine interface {
	Group(relativePath string) RouterGroup
	FormatBindingError(bindError error) error
	Run(addr string) error
}

type RouterGroup interface {
	POST(relativePath string, handler HandlerFunc)
	GET(relativePath string, handler HandlerFunc)
}

type RequestContext interface {
	ShouldBindBodyWith(targetObjectPointer interface{}, bindingBody binding.BindingBody) (err error)
	AbortWithStatusJSON(statusCode int, responseObject interface{})
	JSON(code int, obj interface{})
}