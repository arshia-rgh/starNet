package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/internal/services"
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

// inject user service to user controller

func NewUserController(userService services.UserService) UserController {
	return &userController{userService: userService}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	err := ctx.JSON("succeed")
	if err != nil {
		return err
	}
	return nil
	//make dto
}
