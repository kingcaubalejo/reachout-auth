package service

import (
	"reachout-auth/model"
	"reachout-auth/repository"
)

func ReadAllUser() []*model.User {
	return repository.GetUsers()
}

func ReadUser(id string) *model.User {
	return repository.FindUserById(id)
}

func SaveUser(user *model.User) *model.User {
	return repository.Save(user)
}