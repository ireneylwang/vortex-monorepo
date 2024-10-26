package router

import (
	"export/internal/adapter/controller"
	"github.com/gin-gonic/gin"
)

func ArchiveRouters(rootGroup *gin.RouterGroup) {
	rootGroup.GET("/archives", controller.QueryArchives)
	rootGroup.POST("/archives", controller.CreateArchive)
	rootGroup.PUT("/archives/:id", controller.UpdateArchive)
	rootGroup.DELETE("/archives/:id", controller.DeleteArchive)
}
