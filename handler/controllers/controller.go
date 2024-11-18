package controllers

import "golang_template/internal/services"

type Controllers interface {
	UserController() UserController
	VideoController() VideoController
}

type controllers struct {
	userController  UserController
	videoController VideoController
}

func NewControllers(service services.Service) Controllers {
	userController := NewUserController(service.UserService())
	videoController := NewVideoController(service.VideoService())
	return &controllers{userController: userController, videoController: videoController}
}

func (c *controllers) UserController() UserController {
	return c.userController
}
func (c *controllers) VideoController() VideoController { return c.videoController }
