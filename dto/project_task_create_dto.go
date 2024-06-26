package dto

// ProjectTaskCreateDTO info
// @Description ProjectTask information create dto
type ProjectTaskCreateDTO struct {
	Project         string `json:"project"`
	Manager         string `json:"manager,omitempty"`
	Department      string `json:"department,omitempty"`
	TaskDescription string `json:"task_description,omitempty"`
	Status          string `json:"status"`
} //@name ProjectTaskCreateDTO
