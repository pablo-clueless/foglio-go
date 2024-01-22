package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Team        *Team              `json:"team" bson:"team"`
	PostedBy    *User              `json:"posted_by" bson:"posted_by"`
	PostedOn    primitive.DateTime `json:"posted_on" bson:"posted_on"`
	Deadline    primitive.DateTime `json:"deadline" bson:"deadline"`
	Location    string             `json:"location" bson:"location"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Tags        []string           `json:"tags" bson:"tags"`
	Url         string             `json:"url" bson:"url"`
}
