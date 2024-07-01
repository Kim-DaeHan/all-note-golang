package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type MeetingService interface {
	// GetAllTodo() ([]models.Todo, error)
	GetMeeting(id string) (*models.Meeting, error)
	// GetTodoByUser(userId string) ([]models.Todo, error)
	CreateMeeting(dto *dto.MeetingCreateDTO) error
	// UpdateTodo(id string, dto *dto.TodoUpdateDTO) (*models.Todo, error)
	// DeleteTodo(id string) error
}
