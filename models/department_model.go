package models

type Department struct {
	Name  string `bson:"name" json:"name"`
	Depth string `bson:"depth" json:"depth"`
}
