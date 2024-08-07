package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/Kim-DaeHan/all-note-golang/services"
	"go.mongodb.org/mongo-driver/mongo"
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

	// note
	noteCollection *mongo.Collection
	noteService    services.NoteService
	noteHandler    handlers.NoteHandler
	noteRoute      NoteRoutes

	// todo
	todoCollection *mongo.Collection
	todoService    services.TodoService
	todoHandler    handlers.TodoHandler
	todoRoute      TodoRoutes

	// project
	projectCollection *mongo.Collection
	projectService    services.ProjectService
	projectHandler    handlers.ProjectHandler
	projectRoute      ProjectRoutes

	// project-task
	projectTaskCollection *mongo.Collection
	projectTaskService    services.ProjectTaskService
	projectTaskHandler    handlers.ProjectTaskHandler
	projectTaskRoute      ProjectTaskRoutes

	// meeting
	meetingCollection *mongo.Collection
	meetingService    services.MeetingService
	meetingHandler    handlers.MeetingHandler
	meetingRoute      MeetingRoutes

	// job-application
	jobApplicationCollection *mongo.Collection
	jobApplicationService    services.JobApplicationService
	jobApplicationHandler    handlers.JobApplicationHandler
	jobApplicationRoute      JobApplicationRoutes
)
