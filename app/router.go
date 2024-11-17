package app

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/handler/controllers"
	"golang_template/handler/middlewares"
	"golang_template/handler/routers"
)

func (a *application) InitRouter(app *fiber.App, controller controllers.Controllers) routers.Router {
	router := routers.NewRouter(controller)
	middleware := middlewares.NewMiddleware()

	routeGroup := app.Group("/api")
	protectedRoutes := routeGroup.Group("/protected")
	publicRoutes := routeGroup.Group("/public")

	protectedRoutes.Use(middleware.Auth().Handle())

	router.AddPublicRoutes(publicRoutes)
	router.AddProtectedRoutes(protectedRoutes)

	return router
}
