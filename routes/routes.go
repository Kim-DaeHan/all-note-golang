package routes

import (
	"context"

	"github.com/Kim-DaeHan/all-note-golang/database"
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/Kim-DaeHan/all-note-golang/services/impl"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")

	userRoute.SetUserRoutes(apiGroup)
	authRoute.SetAuthRoutes(apiGroup, userCollection)
	noteRoute.SetNoteRoutes(apiGroup)
	todoRoute.SetTodoRoutes(apiGroup)
	projectRoute.SetProjectRoutes(apiGroup)
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

	// department
	// departmentCollection = database.GetCollection(db, "departments")

	// note
	noteCollection = database.GetCollection(db, "notes")
	noteService = impl.NewNoteServiceImpl(noteCollection)
	noteHandler = handlers.NewNoteController(noteService)
	noteRoute = NewNoteRoutes(noteHandler)

	// todo
	todoCollection = database.GetCollection(db, "todos")
	todoService = impl.NewTodoServiceImpl(todoCollection)
	todoHandler = handlers.NewTodoController(todoService)
	todoRoute = NewTodoRoutes(todoHandler)

	// project
	projectCollection = database.GetCollection(db, "projects")
	projectService = impl.NewProjectServiceImpl(projectCollection)
	projectHandler = handlers.NewProjectController(projectService)
	projectRoute = NewProjectRoutes(projectHandler)
}
