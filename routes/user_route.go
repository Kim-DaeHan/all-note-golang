package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.GET("/", handlers.GetAllUser)
		users.POST("/", handlers.CreateUser)
	}
}
