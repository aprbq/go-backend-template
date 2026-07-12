package service

import "go-backend-template/internal/dto"

type UserService interface {
	GetProfile(userID uint) (*dto.UserResponse, error)
	UpdateProfile(userID uint, input dto.UserRequest) (*dto.UserResponse, error)
	DeleteUser(userID uint) error
	GetAllUsers() ([]dto.UserResponse, error)
}
