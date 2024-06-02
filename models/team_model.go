package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Team info
// @Description Team information
type Team struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Department Department         `bson:"department,omitempty" json:"department,omitempty"`
} //@name Team
