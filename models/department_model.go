package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Department struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Depth string             `bson:"depth" json:"depth"`
}
