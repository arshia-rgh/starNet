package controllers

import "golang_template/internal/services"

type Controllers interface {
	UserController() UserController
}

type controllers struct {
	userController UserController
}

func NewControllers(service services.Service) Controllers {
	userController := NewUserController(service.UserService())
	return &controllers{userController: userController}
}

func (c *controllers) UserController() UserController {
	return c.userController
}
