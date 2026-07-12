package repository

import "go-backend-template/internal/model"

type UserRepository interface {
	Create(user *model.User) error
	FetchAll() ([]model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	Update(user *model.User) error
	Delete(user *model.User) error
}
