package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Education struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Institution string             `json:"institution" bson:"institution"`
	Url         string             `json:"url" bson:"url"`
	Degree      string             `json:"degree" bson:"degree"`
	Major       string             `json:"major" bson:"major"`
	Location    string             `json:"location" bson:"location"`
	From        primitive.DateTime `json:"from" bson:"from"`
	To          primitive.DateTime `json:"to" bson:"to"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
}
