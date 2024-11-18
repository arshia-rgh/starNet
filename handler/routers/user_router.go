package routers

import (
	"golang_template/handler/controllers"

	"github.com/gofiber/fiber/v2"
)

type UserRouter interface {
	AddProtectedRoutes(router fiber.Router)
	AddPublicRoutes(router fiber.Router)
}

type userRouter struct {
	Controller controllers.UserController
}

func NewUserRouter(userController controllers.UserController) UserRouter {
	return &userRouter{Controller: userController}
}

func (r userRouter) AddProtectedRoutes(router fiber.Router) {

}

func (r userRouter) AddPublicRoutes(router fiber.Router) {
	router.Post("/login", r.Controller.Login)
	router.Post("/register", r.Controller.Register)
}
