package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Feature struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Year        primitive.DateTime `json:"year" bson:"year"`
	Publisher   string             `json:"publisher" bson:"publisher"`
	Url         string             `json:"url" bson:"url"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
}
