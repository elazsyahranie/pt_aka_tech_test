package users

import (
	"errors"
	"pt_aka_tech_test/helpers"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user PostBody) (User, error)
	Login (logIn LogInRequest) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(userRequest PostBody) (User, error) {
	var id = uuid.New()
	duplicate, err := s.repository.FindByEmailLike(userRequest.Email)
	if err == nil {
		if duplicate.ID != id {
			return User{}, errors.New(helpers.DUPLICATE)
		}
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 14)
	var hashedPassword = string(bytes)

	user := User{
		ID:        id,
		Name: 	   userRequest.Name,
		Password:  hashedPassword,
		Email:     userRequest.Email,
	}

	newUser, err := s.repository.Register(user)
	return newUser, err
}

func (s *service) GenerateToken(dataForToken DataForToken) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": dataForToken.ID,
		"email": dataForToken.Email,                  
		"exp": time.Now().Add(24 * time.Hour).Unix(), 		// Expiration time
		"iat": time.Now().Unix(),                 		    // Issued at
	})
	tokenString, err := claims.SignedString(helpers.SECRET_KEY)
    if err != nil {
        return "", err
    }

	return tokenString, nil
}

func (s *service) Login(logInRequest LogInRequest) (string, error) {
	user, err := s.repository.FindByEmailLike(logInRequest.Email)
	if err != nil {
		return "", errors.New(helpers.UNAUTHORIZED)
	}

	// Check if password inserted is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(logInRequest.Password))
	if err != nil {
		return "", errors.New(helpers.UNAUTHORIZED)
	}

	tokenRequest := DataForToken{
		ID: user.ID,
		Email: user.Email,
	}

	token, err := s.GenerateToken(tokenRequest)

	return token, err
}
