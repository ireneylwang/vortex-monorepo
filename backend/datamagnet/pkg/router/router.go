package router

import (
	"datamagnet/internal/adapter/controller"
	"github.com/gin-gonic/gin"
)

func IntegrationRouters(rootGroup *gin.RouterGroup) {
	rootGroup.POST("/data-magnet/integrations", controller.AddIntegration)
	rootGroup.DELETE("/data-magnet/integrations/:id", controller.RemoveIntegration)
}
