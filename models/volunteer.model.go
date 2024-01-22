package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Volunteering struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Organization string             `json:"organization" bson:"organization"`
	Location     string             `json:"location" bson:"location"`
	From         primitive.DateTime `json:"from" bson:"from"`
	To           primitive.DateTime `json:"to" bson:"to"`
	Url          string             `json:"url" bson:"url"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	Images       []string           `json:"images,omitempty" bson:"images,omitempty"`
}
