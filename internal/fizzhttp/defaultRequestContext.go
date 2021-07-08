package fizzhttp

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

func NewDefaultContext(ginContext *gin.Context) *defaultRequestContext {
	return &defaultRequestContext{
		internalContext: ginContext,
	}
}

func (context *defaultRequestContext) ShouldBindBodyWithJSON(targetObjectPointer interface{}) (err error) {
	return context.internalContext.ShouldBindBodyWith(targetObjectPointer, binding.JSON)
}

func (context *defaultRequestContext) AbortWithStatusJSON(statusCode int, responseObject interface{}) {
	context.internalContext.AbortWithStatusJSON(statusCode, responseObject)
}

func (context *defaultRequestContext) SendJSONResponse(code int, obj interface{}) {
	context.internalContext.JSON(code, obj)
}

func (context *defaultRequestContext) GetResponseWriter() http.ResponseWriter {
	return context.internalContext.Writer
}

func (context *defaultRequestContext) GetRequest() *http.Request {
	return context.internalContext.Request
}

func handleGinRequest(handler HandlerFunc, context *gin.Context) {
	requestContext := NewDefaultContext(context)
	handler(requestContext)
}
