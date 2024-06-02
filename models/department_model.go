package models

// Department info
// @Description Department information
type Department struct {
	Name  string `bson:"name" json:"name"`
	Depth string `bson:"depth" json:"depth"`
} //@name Department
