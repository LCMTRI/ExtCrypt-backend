package models

import "time"

type User struct {
	UserId        string    `json:"user_id" bson:"user_id"`
	Name          string    `json:"name" bson:"name"`
	Email         string    `json:"email" bson:"email"`
	Role          string    `json:"role" bson:"role"`
	RememberToken string    `json:"remember_token" bson:"remember_token"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}
