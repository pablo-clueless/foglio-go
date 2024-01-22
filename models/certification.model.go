package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Certification struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IssuedBy     string             `json:"issued_by" bson:"issued_by"`
	ExpiryDate   primitive.DateTime `json:"expiry_date" bson:"expiry_date"`
	Name         string             `json:"name" bson:"name"`
	Organization string             `json:"organization" bson:"organization"`
	Year         primitive.DateTime `json:"year" bson:"year"`
	Url          string             `json:"url" bson:"url"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	Images       []string           `json:"images,omitempty" bson:"images,omitempty"`
}
