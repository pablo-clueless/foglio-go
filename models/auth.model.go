package models

type Auth struct {
	Email string `json:"email" bson:"email"`
}
