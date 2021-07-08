package fizzhttp

import "github.com/gin-gonic/gin"

type defaultEngineFactory struct{}

type defaultEngine struct {
	internalEngine *gin.Engine
}

type defaultRouterGroup struct {
	internalRouterGroup *gin.RouterGroup
}

type httpErrorResponse struct {
	Type        string                `json:"type"`
	FieldErrors []*responseFieldError `json:"fieldErrors,omitempty"`
}

type responseFieldError struct {
	Type   string `json:"type"`
	Field  string `json:"field"`
	Detail string `json:"detail"`
}
