package fizzhttp

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var (
	JsonBinding = binding.JSON
	ReleaseMode = gin.ReleaseMode
	TestMode    = gin.TestMode
	DebugMode   = gin.DebugMode
)
