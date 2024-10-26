package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/adapter/controller"
)

type Option func(rootGroup *gin.RouterGroup)

func ServerRouters(rootPath string, options ...Option) *gin.Engine {
	engine := gin.Default()
	rootGroup := engine.Group(rootPath)

	for _, option := range options {
		option(rootGroup)
	}

	return engine
}

func ServerMetricsRouters(rootGroup *gin.RouterGroup) {
	rootGroup.GET("/version", controller.Version)
}
