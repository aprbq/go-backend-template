package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID       primitive.ObjectID `json:"id"`
	FullName string             `json:"full_name"`
	Email    string             `json:"email"`
	Picture  string             `json:"picture"`
}

type UserRequest struct {
	Firstname string `json:"first_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
}
