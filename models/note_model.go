package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Note info
// @Description Note information
type Note struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Author     primitive.ObjectID `bson:"author" json:"author"`
	AuthorInfo []User             `bson:"author_info,omitempty" json:"author_info,omitempty"`
	Text       string             `bson:"text" json:"text"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
} //@name Note
