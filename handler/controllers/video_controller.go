package controllers

import "golang_template/internal/services"

type VideoController interface {
}

type videoController struct {
	videoService services.VideoService
}

func NewVideoController(videoService services.VideoService) VideoController {
	return &videoController{videoService: videoService}
}
