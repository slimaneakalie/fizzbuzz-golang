package fizzhttp

import "github.com/gin-gonic/gin"

type mainEngineFactory struct{}

type mainEngine struct {
	internalEngine *gin.Engine
}

type mainRouterGroup struct {
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
