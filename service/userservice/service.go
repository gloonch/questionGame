package userservice

import (
	"fmt"
	"questionGame/entity"
	"questionGame/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO: we should verify phone number by verification code

	// validate phone number
	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("invalid phone number")
	}

	// check uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error while checking if phone number exists: %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number %s is already used", req.PhoneNumber)
		}
	}

	// validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name should be greater than 3 characters")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	// create new user in storage
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error while registering user: %w", err)
	}

	// return created user
	return RegisterResponse{
		User: createdUser,
	}, nil
}
