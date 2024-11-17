package app

import (
	"golang_template/handler/controllers"
	"golang_template/handler/middlewares"
	"golang_template/handler/routers"
	"golang_template/internal/config"

	"github.com/gofiber/fiber/v2"
)

func (a *application) InitRouter(app *fiber.App, controller controllers.Controllers, cfg *config.JWTConfig) routers.Router {
	router := routers.NewRouter(controller)
	middleware := middlewares.NewMiddleware(cfg)

	routeGroup := app.Group("/api")
	protectedRoutes := routeGroup.Group("/protected")
	publicRoutes := routeGroup.Group("/public")

	protectedRoutes.Use(middleware.Auth())

	router.AddPublicRoutes(publicRoutes)
	router.AddProtectedRoutes(protectedRoutes)

	return router
}
