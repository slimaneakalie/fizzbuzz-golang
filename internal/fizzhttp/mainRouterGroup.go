package fizzhttp

import "github.com/gin-gonic/gin"

func (group *mainRouterGroup) POST(relativePath string, handler HandlerFunc) {
	group.internalRouterGroup.POST(relativePath, func(context *gin.Context) {
		handler(context)
	})
}

func (group *mainRouterGroup) GET(relativePath string, handler HandlerFunc) {
	group.internalRouterGroup.GET(relativePath, func(context *gin.Context) {
		handler(context)
	})
}
