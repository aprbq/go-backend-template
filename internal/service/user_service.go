package service

import (
	"errors"
	"go-backend-template/internal/dto"
	"go-backend-template/internal/errs"
	"go-backend-template/internal/repository"
	"time"
)

// Business logic ของ user
type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) GetProfile(userID uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	u := dto.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Picture:  user.Picture,
	}
	return &u, nil
}

func (s userService) UpdateProfile(userID uint, input dto.UserRequest) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if input.Email != "" && input.Email != user.Email {
		existing, err := s.userRepo.FindByEmail(input.Email)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			return nil, errors.New("email already in use")
		}
		user.Email = input.Email
	}

	if input.Firstname != "" {
		user.Firstname = input.Firstname
	}

	if input.Surname != "" {
		user.Surname = input.Surname
	}

	user.UpdatedAt = time.Now().UTC()
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	u := dto.UserResponse{
		ID: user.ID,
	}

	return &u, nil
}

func (s userService) DeleteUser(userID uint) error {
	return nil
}

func (s userService) GetAllUsers() ([]dto.UserResponse, error) {

	users, err := s.userRepo.FetchAll()
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	responses := make([]dto.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, dto.UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
			Picture:  user.Picture,
		})
	}

	return responses, nil
}
