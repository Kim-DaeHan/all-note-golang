package dto

// ProjectTaskUpdateDTO info
// @Description ProjectTask information update dto
type ProjectTaskUpdateDTO struct {
	Manager         string `json:"manager,omitempty"`
	Department      string `json:"department,omitempty"`
	TaskDescription string `json:"task_description,omitempty"`
	Status          string `json:"status"`
} //@name ProjectTaskUpdateDTO
