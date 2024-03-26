package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler handlers.UserHandler
}

func NewUserRoutes(userHandler handlers.UserHandler) UserRoutes {
	return UserRoutes{userHandler}
}

func (ur *UserRoutes) SetUserRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")

	users.GET("/", ur.userHandler.GetAllUser)
	users.POST("/", ur.userHandler.CreateUser)
	users.POST("/upsert", ur.userHandler.UpsertUser)

}
