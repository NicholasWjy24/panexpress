package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid username or password")

var jwtSecret = []byte("your_super_secret_key_change_this_in_production")

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

func AuthService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Register(input RegisterInput) (*User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := &User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword),
		RoleLevel: 3,
	}

	if err := s.repository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(input LoginInput) (*User, string, error) {

	user, err := s.repository.FindUserByUsername(input.Username)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	claims := jwt.MapClaims{
		"id":         user.ID,
		"role_level": user.RoleLevel,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}
