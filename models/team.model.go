package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Manager   *User              `json:"manager" bson:"manager"`
	Members   []*User            `json:"members,omitempty" bson:"members,omitempty"`
	Location  string             `json:"location" bson:"location"`
	Type      TeamType           `json:"type" bson:"type"`
	Url       string             `json:"url" bson:"url"`
	Logo      string             `json:"logo" bson:"logo"`
	Images    []string           `json:"images,omitempty" bson:"images,omitempty"`
	Verified  bool               `json:"verified" bson:"verified"`
	OpenRoles []*Job             `json:"open_roles,omitempty" bson:"open_roles,omitempty"`
	Mission   string             `json:"mission,omitempty" bson:"mission,omitempty"`
}

type TeamType string

const (
	TeamTypeRemote TeamType = "remote"
	TeamTypeOnsite TeamType = "onsite"
	TeamTypeHybrid TeamType = "hybrid"
)
