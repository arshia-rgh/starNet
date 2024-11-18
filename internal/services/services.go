package services

import "golang_template/internal/repositories"

type Service interface {
	UserService() UserService
	VideoService() VideoService
}

type service struct {
	userService  UserService
	videoService VideoService
}

func NewService(repo repositories.Repository) Service {
	userService := NewUserService(repo.UserRepository())
	return &service{userService: userService}
}

func (s *service) UserService() UserService   { return s.userService }
func (s *service) VideoService() VideoService { return s.videoService }
