package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type NoteService interface {
	GetAllNote() ([]models.Note, error)
	GetNote(id string) (*models.Note, error)
	GetNoteByUser(userId string) ([]models.Note, error)
	CreateNote(dto *dto.NoteCreateDTO) error
	UpdateNote(id string, dto *dto.NoteUpdateDTO) (*models.Note, error)
	DeleteNote(id string) error
}
