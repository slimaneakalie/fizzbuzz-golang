package fizzhttp

import (
	"github.com/gin-gonic/gin"
)

func (group *defaultRouterGroup) POST(relativePath string, handler HandlerFunc) {
	group.internalRouterGroup.POST(relativePath, func(context *gin.Context) {
		requestContext := NewDefaultContext(context)
		handler(requestContext)
	})
}

func (group *defaultRouterGroup) GET(relativePath string, handler HandlerFunc) {
	group.internalRouterGroup.GET(relativePath, func(context *gin.Context) {
		requestContext := NewDefaultContext(context)
		handler(requestContext)
	})
}
