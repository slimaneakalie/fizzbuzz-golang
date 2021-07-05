package fizzhttp

import "github.com/gin-gonic/gin"

type defaultEngineFactory struct{}

type defaultEngine struct {
	internalEngine *gin.Engine
}

type defaultRouterGroup struct {
	internalRouterGroup *gin.RouterGroup
}

type badRequestServerResponsePayload struct {
	Type     string       `json:"type"`
	Metadata []fieldError `json:"errors,omitempty"`
}

type fieldError struct {
	Type   string `json:"type"`
	Field  string `json:"field"`
	Detail string `json:"detail"`
}
