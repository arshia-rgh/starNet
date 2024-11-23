package controllers

import (
	"errors"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"github.com/gofiber/fiber/v2"
	"golang_template/internal/services"
	"golang_template/internal/services/dto"
	"golang_template/pkg"
	"log"
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type userController struct {
	userService services.UserService
}

// inject user service to user controller

func NewUserController(userService services.UserService) UserController {
	return &userController{userService: userService}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	var user dto.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	token, err := c.userService.Login(ctx, user)
	if err != nil {
		log.Println(err)
		if shared.IsNotFound(err) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user with this username and password doesnt exist"})
		}
		if errors.Is(err, pkg.ErrInvalidToken) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "token is invalid"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}

func (c *userController) Register(ctx *fiber.Ctx) error {
	var user dto.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "wrong data format"})
	}
	dbUser, err := c.userService.Register(ctx, user)
	if err != nil {
		if shared.IsArangoErrorWithCode(err, 409) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "user with this username already exists"})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to create new user"})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dbUser)
}
