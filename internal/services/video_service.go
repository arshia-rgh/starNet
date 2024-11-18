package services

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/internal/ent"
	"golang_template/internal/repositories"
)

type VideoService interface {
	CreateVideo(ctx fiber.Ctx) (*ent.Video, error)
}

type videoService struct {
	repo repositories.VideoRepository
}
