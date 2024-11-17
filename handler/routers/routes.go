package routers

import (
	"golang_template/handler/controllers"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	AddProtectedRoutes(router fiber.Router)
	AddPublicRoutes(router fiber.Router)
}

type router struct {
	userRouter UserRouter
}

func NewRouter(controllers controllers.Controllers) Router {
	userRouter := NewUserRouter(controllers.UserController())
	return &router{userRouter: userRouter}
}

// protected routes means protected by auth (logged in needed)

func (r router) AddProtectedRoutes(router fiber.Router) {

}

func (r router) AddPublicRoutes(router fiber.Router) {
	r.userRouter.AddPublicRoutes(router)
}
