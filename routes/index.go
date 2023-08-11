package routes

import (
	"Go-project/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine){
	router.POST("/api/login",controllers.Login)
	// router.POST("/api/register")
	// router.GET("/api/profile:id")
}