package services

import (
	"fmt"
	config2 "golang_template/internal/config"
	"golang_template/internal/ent"
	"golang_template/internal/repositories"
	"golang_template/internal/services/dto"
	"golang_template/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Login(ctx *fiber.Ctx, user dto.User) (string, error)
	Register(ctx *fiber.Ctx, user dto.User) (*ent.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(ctx *fiber.Ctx, user dto.User) (string, error) {
	config, err := config2.LoadConfig("config/config.yml")
	if err != nil {
		return "", err
	}

	storedUser, err := s.repo.Get(ctx.Context(), user)
	if err != nil {
		return "", err
	}
	if !pkg.VerifyPassword(user.Password, storedUser.Password) {
		return "", fmt.Errorf("password is wrong")
	}

	token, err := pkg.GenerateToken(storedUser.ID, string(storedUser.Role), config.JWT.Secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Register(ctx *fiber.Ctx, user dto.User) (*ent.User, error) {
	hashedPass, _ := pkg.HashPassword(user.Password)
	user.Password = hashedPass
	return s.repo.CreateUser(ctx.Context(), user)

}
