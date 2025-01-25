package user

import (
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(userRequest UserRequest) (User, error)
	Login(userLogin UserLogin) (User, error)
	CheckUser(id string) (User, error)
	GetByID(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByID(id int) (User, error) {
	return s.repository.GetByID(id)
}

func (s *service) Register(userRequest UserRequest) (User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 14)

	println(userRequest.Password)
	user := User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: password,
	}
	return s.repository.Create(user)
}

func (s *service) Login(userLogin UserLogin) (User, error) {
	user, err := s.repository.GetByEmail(userLogin.Email)
	if err != nil {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userLogin.Password))

	if err != nil {
		println("Enter")
		return user, errors.New("incorrect password")
	}

	return user, err
}

func (s *service) CheckUser(id string) (User, error) {
	userID, _ := strconv.Atoi(id)

	user, err := s.repository.GetByID(userID)
	if err != nil {
		return user, errors.New("user not found")
	}

	return user, err
}
