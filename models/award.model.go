package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Award struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Url         string             `json:"url" bson:"url"`
	PrsentedBy  string             `json:"presented_by" bson:"presented_by"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Year        primitive.DateTime `json:"year" bson:"year"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
}
