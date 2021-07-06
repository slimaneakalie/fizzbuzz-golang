package fizzhttp

import "github.com/gin-gonic/gin"

type defaultEngineFactory struct{}

type defaultEngine struct {
	internalEngine *gin.Engine
}

type defaultRouterGroup struct {
	internalRouterGroup *gin.RouterGroup
}

type httpErrorResponseMetadata struct {
	Type        string        `json:"type"`
	FieldErrors []*fieldError `json:"fieldErrors,omitempty"`
}

type fieldError struct {
	Type   string `json:"type"`
	Field  string `json:"field"`
	Detail string `json:"detail"`
}
