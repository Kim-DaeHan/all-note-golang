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
	projectTaskRoute.SetProjectTaskRoutes(apiGroup)
	meetingRoute.SetMeetingRoutes(apiGroup)
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
	userHandler = handlers.NewUserHandler(userService)
	userRoute = NewUserRoutes(userHandler)

	// auth
	authHandler = handlers.NewAuthHandler(userService)
	authRoute = NewAuthRoutes(authHandler)

	// department
	// departmentCollection = database.GetCollection(db, "departments")

	// note
	noteCollection = database.GetCollection(db, "notes")
	noteService = impl.NewNoteServiceImpl(noteCollection)
	noteHandler = handlers.NewNoteHandler(noteService)
	noteRoute = NewNoteRoutes(noteHandler)

	// todo
	todoCollection = database.GetCollection(db, "todos")
	todoService = impl.NewTodoServiceImpl(todoCollection)
	todoHandler = handlers.NewTodoHandler(todoService)
	todoRoute = NewTodoRoutes(todoHandler)

	// project
	projectCollection = database.GetCollection(db, "projects")
	projectService = impl.NewProjectServiceImpl(projectCollection)
	projectHandler = handlers.NewProjectHandler(projectService)
	projectRoute = NewProjectRoutes(projectHandler)

	// project-tasks
	projectTaskCollection = database.GetCollection(db, "project_tasks")
	projectTaskService = impl.NewProjectTaskServiceImpl(projectTaskCollection)
	projectTaskHandler = handlers.NewProjectTaskHandler(projectTaskService)
	projectTaskRoute = NewProjectTaskRoutes(projectTaskHandler)

	// meeting
	meetingCollection = database.GetCollection(db, "meetings")
	meetingService = impl.NewMeetingServiceImpl(meetingCollection)
	meetingHandler = handlers.NewMeetingHandler(meetingService)
	meetingRoute = NewMeetingRoutes(meetingHandler)
}
