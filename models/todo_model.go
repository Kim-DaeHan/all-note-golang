package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Todo info
// @Description Todo information
type Todo struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Task           string             `bson:"task" json:"task"`
	Status         string             `bson:"status" json:"status"`
	Project        primitive.ObjectID `bson:"project" json:"project"`
	ProjectrInfo   []Project          `bson:"project_info" json:"project_info"`
	StartDt        time.Time          `bson:"start_dt" json:"start_dt"`
	EndDt          time.Time          `bson:"end_dt" json:"end_dt"`
	User           primitive.ObjectID `bson:"user" json:"user"`
	UserInfo       []User             `bson:"user_info" json:"user_info"`
	Department     primitive.ObjectID `bson:"department,omitempty" json:"department,omitempty"`
	DepartmentInfo []Department       `bson:"department_info,omitempty" json:"department_info,omitempty"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
} //@name Todo
