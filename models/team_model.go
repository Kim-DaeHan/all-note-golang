package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Department primitive.ObjectID `bson:"department" json:"department"`
}
