package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Username       string              `json:"username" bson:"username" unique:"true"`
	Name           string              `json:"name" bson:"name"`
	Role           string              `json:"role" bson:"role"`
	Location       string              `json:"location" bson:"location"`
	About          string              `json:"about" bson:"about"`
	Website        string              `json:"website" bson:"website"`
	Email          string              `json:"email" bson:"email" unique:"true"`
	Image          string              `json:"image,omitempty" bson:"image,omitempty"`
	Contact        map[string]string   `json:"contact" bson:"contact"`
	Password       string              `json:"password" bson:"password"`
	Createdat      primitive.Timestamp `json:"created_at" bson:"created_at"`
	Projects       []*Project          `json:"projects,omitempty" bson:"projects,omitempty"`
	SideProjects   []*Project          `json:"side_projects,omitempty" bson:"side_projects,omitempty"`
	Education      []*Education        `json:"education,omitempty" bson:"education,omitempty"`
	Experience     []*Experience       `json:"experience,omitempty" bson:"experience,omitempty"`
	Awards         []*Award            `json:"awards,omitempty" bson:"awards,omitempty"`
	Volunteering   []*Volunteering     `json:"volunteering,omitempty" bson:"volunteering,omitempty"`
	Speaking       []*Speaking         `json:"speaking,omitempty" bson:"speaking,omitempty"`
	Exhibitions    []*Exhibition       `json:"exhibitions,omitempty" bson:"exhibitions,omitempty"`
	Certifications []*Certification    `json:"certifications,omitempty" bson:"certifications,omitempty"`
	Features       []*Feature          `json:"features,omitempty" bson:"features,omitempty"`
}
