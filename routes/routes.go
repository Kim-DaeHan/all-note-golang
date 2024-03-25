package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/database"
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection *mongo.Collection
	userService    services.UserService
	userHandler    handlers.UserHandler
	userRoute      UserRoutes
)

func SetupRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")

	userRoute.SetUserRoutes(apiGroup)

}

func SetDependency(db *mongo.Client) {
	// user
	userCollection = database.GetCollection(db, "users")
	userService = services.NewUserServiceImpl(userCollection)
	userHandler = handlers.NewUserController(userService)
	userRoute = NewUserRoutes(userHandler)

}
