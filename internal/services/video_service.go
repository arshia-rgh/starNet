package services

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/internal/ent"
	"golang_template/internal/repositories"
	"golang_template/internal/services/dto"
)

type VideoService interface {
	CreateVideo(ctx *fiber.Ctx, video dto.Video) (*ent.Video, error)
	GetAllVideos(ctx *fiber.Ctx) ([]*ent.Video, error)
	GetVideoByTitle(ctx *fiber.Ctx, video dto.Video) (*ent.Video, error)
}

type videoService struct {
	repo repositories.VideoRepository
}

func NewVideoService(repo repositories.VideoRepository) VideoService {
	return &videoService{repo: repo}
}

func (v *videoService) CreateVideo(ctx *fiber.Ctx, video dto.Video) (*ent.Video, error) {
	return v.repo.CreateVideo(ctx.Context(), video)
}
func (v *videoService) GetAllVideos(ctx *fiber.Ctx) ([]*ent.Video, error) {
	return v.repo.GetAllVideos(ctx.Context())
}
func (v *videoService) GetVideoByTitle(ctx *fiber.Ctx, video dto.Video) (*ent.Video, error) {
	return v.repo.GetVideoByTitle(ctx.Context(), video)
}
