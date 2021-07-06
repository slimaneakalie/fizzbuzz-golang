package fizzhttp

type EngineFactory interface {
	NewEngine() Engine
}

type Engine interface {
	Group(relativePath string) RouterGroup
	FormatBindingError(bindError error) error
	Run(port int) error
}

type RouterGroup interface {
	POST(relativePath string, handler HandlerFunc)
	GET(relativePath string, handler HandlerFunc)
}

type RequestContext interface {
	ShouldBindBodyWith(targetObjectPointer interface{}, requestBody bindingBodyType) (err error)
	AbortWithStatusJSON(statusCode int, responseObject interface{})
	JSON(code int, obj interface{})
}
