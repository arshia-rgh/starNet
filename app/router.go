package app

import (
	"golang_template/handler/controllers"
	"golang_template/handler/middlewares"
	"golang_template/handler/routers"

	"github.com/gofiber/fiber/v2"
)

func (a *application) InitRouter(app *fiber.App, controller controllers.Controllers) routers.Router {
	router := routers.NewRouter(controller)
	middleware := middlewares.NewMiddleware()

	routeGroup := app.Group("/api")
	protectedRoutes := setupProtectedRoutes(routeGroup.Group("/protected"), middleware)
	publicRoutes := setupPublicRoutes(routeGroup.Group("/public"))

	router.AddPublicRoutes(publicRoutes)
	router.AddProtectedRoutes(protectedRoutes)

	return router
}

func setupProtectedRoutes(group fiber.Router, middleware middlewares.Middleware) fiber.Router {
	group.Use(middleware.Auth().Handle())
	group.Use(middleware.Authorization().Handle())
	return group
}

func setupPublicRoutes(group fiber.Router) fiber.Router {
	// any additional logic or middlewares if needed
	return group
}
