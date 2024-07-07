package dto

import "time"

// TodoUpdateDTO info
// @Description Todo information update dto
type TodoUpdateDTO struct {
	Task       string    `json:"task,omitempty"`
	Status     string    `json:"status,omitempty"`
	Project    string    `json:"project,omitempty"`
	StartDt    time.Time `json:"start_dt,omitempty"`
	EndDt      time.Time `json:"end_dt,omitempty"`
	Department string    `json:"department,omitempty"`
} //@name TodoUpdateDTO
