package routers

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/handler/controllers"
)

type VideoRouter interface {
	AddProtectedRoutes(router fiber.Router)
	AddPublicRoutes(router fiber.Router)
}

type videoRouter struct {
	Controller controllers.VideoController
}

func NewVideoRouter(controller controllers.VideoController) VideoRouter {
	return &videoRouter{Controller: controller}
}

func (v *videoRouter) AddProtectedRoutes(router fiber.Router) {
	router.Post("/upload-video", v.Controller.UploadVideo)
}
func (v *videoRouter) AddPublicRoutes(router fiber.Router) {
	router.Get("/videos", v.Controller.ShowAllVideos)
}
