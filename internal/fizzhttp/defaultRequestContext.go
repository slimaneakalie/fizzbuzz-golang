package fizzhttp

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

func NewDefaultContext(ginContext *gin.Context) *defaultRequestContext {
	return &defaultRequestContext{
		internalContext: ginContext,
	}
}

func (context *defaultRequestContext) ShouldBindBodyWithJSON(targetObjectPointer interface{}) error {
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

func (context *defaultRequestContext) GetResponseStatus() int {
	return context.internalContext.Writer.Status()
}

func (context *defaultRequestContext) GetJsonStringQuery() string {
	var requestBody interface{}
	context.internalContext.ShouldBindBodyWith(&requestBody, binding.JSON)
	jsonBytes, _ := json.Marshal(requestBody)
	return string(jsonBytes)
}

func (context *defaultRequestContext) Next() {
	context.internalContext.Next()
}

func handleGinRequest(handler HandlerFunc, context *gin.Context) {
	requestContext := NewDefaultContext(context)
	handler(requestContext)
}
