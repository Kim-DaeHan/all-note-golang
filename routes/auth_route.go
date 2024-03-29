package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/Kim-DaeHan/all-note-golang/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRoutes struct {
	authHandler handlers.AuthHandler
}

func NewAuthRoutes(authHandler handlers.AuthHandler) AuthRoutes {
	return AuthRoutes{authHandler}
}

func (ar *AuthRoutes) SetAuthRoutes(router *gin.RouterGroup, collection *mongo.Collection) {
	auths := router.Group("/auth")

	auths.GET("/google", ar.authHandler.GoogleOAuth)
	auths.GET("/logout", ar.authHandler.LogoutUser)
	auths.GET("/users", middleware.DeserializeUser(collection), ar.authHandler.GetMe)
	auths.GET("/refresh", ar.authHandler.RefreshAccessToken)
}
