package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(input CreateUserInput) (User, error)
	LoginUser(input LoginUserInput) (User, error)
	GetUserById(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByUsername(username string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(input CreateUserInput) (User, error) {
	user := User{}
	user.Username = input.Username
	user.Email = input.Email
	user.Age = input.Age

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByUsername(username string) (User, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}
