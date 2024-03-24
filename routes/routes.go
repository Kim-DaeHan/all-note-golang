package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		UserRoutes(apiGroup)
	}
}
