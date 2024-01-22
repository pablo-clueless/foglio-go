package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Company     string             `json:"company" bson:"company"`
	Url         string             `json:"url" bson:"url"`
	Position    string             `json:"position" bson:"position"`
	From        primitive.DateTime `json:"from" bson:"from"`
	To          primitive.DateTime `json:"to" bson:"to"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Images      []string           `json:"images,omitempty" bson:"images,omitempty"`
}
