package dto

import "time"

// TodoUpdateDTO info
// @Description Todo information update dto
type TodoUpdateDTO struct {
	Task       string    `json:"task"`
	Status     string    `json:"status"`
	Project    string    `json:"project"`
	StartDt    time.Time `json:"start_dt"`
	EndDt      time.Time `json:"end_dt"`
	Department string    `json:"department"`
} //@name TodoUpdateDTO
