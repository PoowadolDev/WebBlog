package service

import (
	"Web_Blog/core/domain"
	"Web_Blog/core/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user domain.User) error
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(user domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	if err = s.repo.Save(user); err != nil {
		return err
	}

	return nil
}
