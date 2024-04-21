package routes

import (
	"context"

	"github.com/Kim-DaeHan/all-note-golang/database"
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/Kim-DaeHan/all-note-golang/services/impl"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// user
	userCollection *mongo.Collection
	userService    services.UserService
	userHandler    handlers.UserHandler
	userRoute      UserRoutes

	// auth
	authHandler handlers.AuthHandler
	authRoute   AuthRoutes
)

func SetupRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")

	userRoute.SetUserRoutes(apiGroup)
	authRoute.SetAuthRoutes(apiGroup, userCollection)

}

func SetDependency(db *mongo.Client) {
	// user
	userCollection = database.GetCollection(db, "users")
	userCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	userService = impl.NewUserServiceImpl(userCollection)
	userHandler = handlers.NewUserController(userService)
	userRoute = NewUserRoutes(userHandler)

	// auth
	authHandler = handlers.NewAuthController(userService)
	authRoute = NewAuthRoutes(authHandler)
}
