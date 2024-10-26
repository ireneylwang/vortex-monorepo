package router

import (
	"github.com/gin-gonic/gin"
	"iam/internal/controller"
)

func AccessRouters(rootGroup *gin.RouterGroup) {
	rootGroup.POST("/login", controller.Login)
}
