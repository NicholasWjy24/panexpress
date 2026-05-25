package auth

import (
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

type Service struct {
	repository *Repository
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Register(input RegisterInput) (*User, error) {
	user := &User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := s.repository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(input LoginInput) (*User, error) {
	user, err := s.repository.FindUserByUsername(input.Username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if user.Password != input.Password {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
