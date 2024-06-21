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
)
