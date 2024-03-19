package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.GET("/", handlers.GetAllUser)
	}
}
