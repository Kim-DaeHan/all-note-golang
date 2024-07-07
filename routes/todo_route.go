package routes

import (
	"github.com/Kim-DaeHan/all-note-golang/handlers"
	"github.com/gin-gonic/gin"
)

type TodoRoutes struct {
	todoHandler handlers.TodoHandler
}

func NewTodoRoutes(todoHandler handlers.TodoHandler) TodoRoutes {
	return TodoRoutes{todoHandler}
}

func (tr *TodoRoutes) SetTodoRoutes(router *gin.RouterGroup) {
	todos := router.Group("/todos")

	todos.GET("/", tr.todoHandler.GetAllTodo)
	todos.GET("/:id", tr.todoHandler.GetTodo)
	todos.GET("/user/:id", tr.todoHandler.GetTodoByUser)
	todos.POST("/", tr.todoHandler.CreateTodo)
	todos.PATCH("/:id", tr.todoHandler.UpdateTodo)
	todos.DELETE("/:id", tr.todoHandler.DeleteTodo)

}
