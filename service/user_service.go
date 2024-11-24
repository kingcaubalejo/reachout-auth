package service

import (
	"reachout-auth/model"
	"reachout-auth/repository"
)

type UserService interface {
	ReadAllUser() ([]*model.User)
	ReadUser(id string) *model.User
	CreateUser(user *model.User) *model.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (s *userService) ReadAllUser() []*model.User {
	return s.userRepository.GetUsers()
}

func (s *userService) ReadUser(id string) *model.User {
	return s.userRepository.FindUserById(id)
}

func (s *userService) CreateUser(user *model.User) *model.User {
	return s.userRepository.Save(user)
}