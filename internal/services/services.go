package services

import "golang_template/internal/repositories"

type Service interface {
}

type service struct {
	userService UserService
}

func NewService(repo repositories.Repository) Service {
	userService := NewUserService(repo.UserRepository())
	return &service{userService: userService}
}
