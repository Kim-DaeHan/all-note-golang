package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProjectTask info
// @Description ProjectTask information
type ProjectTask struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Project         primitive.ObjectID `bson:"project" json:"project"`
	ProjectrInfo    []Project          `bson:"project_info,omitempty" json:"project_info,omitempty"`
	User            primitive.ObjectID `bson:"manager" json:"manager"`
	UserInfo        []User             `bson:"manager_info,omitempty" json:"manager_info,omitempty"`
	Department      primitive.ObjectID `bson:"department,omitempty" json:"department,omitempty"`
	DepartmentInfo  []Department       `bson:"department_info,omitempty" json:"department_info,omitempty"`
	TaskDescription string             `bson:"task_description" json:"task_description"`
	Status          string             `bson:"status" json:"status"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
} //@name ProjectTask
