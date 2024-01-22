package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Speaking struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Event       string             `json:"event" bson:"event"`
	Location    string             `json:"location" bson:"location"`
	Year        primitive.DateTime `json:"year" bson:"year"`
	Url         string             `json:"url" bson:"url"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
}
