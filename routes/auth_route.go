package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	authHandler handlers.AuthHandler
}

func NewAuthRoutes(authHandler handlers.AuthHandler) AuthRoutes {
	return AuthRoutes{authHandler}
}

func (ar *AuthRoutes) SetAuthRoutes(router *gin.RouterGroup) {
	auths := router.Group("/auth")

	auths.GET("/google", ar.authHandler.GoogleOAuth)

}
