package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Url         string             `json:"url" bson:"url"`
	Year        primitive.DateTime `json:"year" bson:"year"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
}
