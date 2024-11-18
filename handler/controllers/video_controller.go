package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/internal/ent"
	"golang_template/internal/services"
	"golang_template/internal/services/dto"
)

type VideoController interface {
	UploadVideo(ctx *fiber.Ctx) error
	ShowAllVideos(ctx *fiber.Ctx) error
}

type videoController struct {
	videoService services.VideoService
}

func NewVideoController(videoService services.VideoService) VideoController {
	return &videoController{videoService: videoService}
}

func (v *videoController) UploadVideo(ctx *fiber.Ctx) error {
	var video dto.Video
	if err := ctx.BodyParser(&video); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid data type"})
	}
	dbVideo, err := v.videoService.CreateVideo(ctx, video)
	if err != nil {
		if ent.IsConstraintError(err) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "video with this title already exists"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(dbVideo)
}

func (v *videoController) ShowAllVideos(ctx *fiber.Ctx) error {
	videos, err := v.videoService.GetAllVideos(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "no videos found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(videos)
}
