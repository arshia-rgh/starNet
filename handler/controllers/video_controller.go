package controllers

import (
	"golang_template/internal/ent"
	"golang_template/internal/services"
	"golang_template/internal/services/dto"
	"io"
	"log"
	"os"
	filepath2 "path/filepath"
	"strconv"

	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"github.com/gofiber/fiber/v2"
)

type VideoController interface {
	UploadVideo(ctx *fiber.Ctx) error
	ShowAllVideos(ctx *fiber.Ctx) error
	PlayVideo(ctx *fiber.Ctx) error
}

type videoController struct {
	videoService services.VideoService
}

func NewVideoController(videoService services.VideoService) VideoController {
	return &videoController{videoService: videoService}
}

// UploadVideo is only accessible by admins
func (v *videoController) UploadVideo(ctx *fiber.Ctx) error {
	uploadDTO, err := parseVideoUploadDTO(ctx)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "failed to parse form values and file"})
	}
	if err := os.MkdirAll("videos", os.ModePerm); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to create upload directory"})
	}

	tempFilePath := filepath2.Join("videos", uploadDTO.Title+"_tmp")

	file, err := uploadDTO.File.Open()
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to open file chunk"})
	}
	defer file.Close()
	tempFile, err := os.OpenFile(tempFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to open temporary file"})
	}
	defer tempFile.Close()

	if _, err := io.Copy(tempFile, file); err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to write file chunk"})
	}
	if uploadDTO.ChunkNumber == uploadDTO.TotalChunk {
		finalFilePath := filepath2.Join("videos", uploadDTO.Title)
		if err := os.Rename(tempFilePath, finalFilePath); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "failed to finalize video file"})
		}

		var video dto.Video
		video.Title = uploadDTO.Title
		video.Description = uploadDTO.Description
		video.FilePath = finalFilePath

		dbVideo, err := v.videoService.CreateVideo(ctx, video)
		if err != nil {
			log.Println(err)
			if shared.IsArangoErrorWithCode(err, 409) {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "video with this title already exists"})
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
		}
		return ctx.Status(fiber.StatusOK).JSON(dbVideo)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "chunk uploaded successfully"})
}

func parseVideoUploadDTO(ctx *fiber.Ctx) (*dto.VideoUpload, error) {
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	chunkNumber, err := strconv.Atoi(ctx.FormValue("chunk_number"))
	if err != nil {
		return nil, err
	}
	totalChunks, err := strconv.Atoi(ctx.FormValue("total_chunks"))
	if err != nil {
		return nil, err
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}

	return &dto.VideoUpload{
		Title:       title,
		Description: description,
		ChunkNumber: chunkNumber,
		TotalChunk:  totalChunks,
		File:        file,
	}, nil
}

// unreleased version of fiber

//func parseVideoUploadDTO(ctx *fiber.Ctx) (*dto.VideoUpload, error) {
//	var video dto.VideoUpload
//	if err := ctx.Bind().Form(&video); err != nil {
//		log.Println(err)
//		return nil, err
//	}
//
//	return &video, nil
//}

func (v *videoController) ShowAllVideos(ctx *fiber.Ctx) error {
	videos, err := v.videoService.GetAllVideos(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "no videos found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(videos)
}

func (v *videoController) PlayVideo(ctx *fiber.Ctx) error {
	var params dto.VideoParams
	err := ctx.ParamsParser(&params)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "wrong params format"})
	}

	video := dto.Video{Title: params.Title}
	dbVideo, err := v.videoService.GetVideoByTitle(ctx, video)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "video not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "server error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(dbVideo.FilePath)

}
