package app

import (
	"golang_template/handler/controllers"
	"golang_template/handler/routers"

	"github.com/gofiber/fiber/v2"
)

func (a *application) InitRouter(app *fiber.App, controller controllers.Controllers) routers.Router {
	router := routers.NewRouter(controller)
	router.AddRoutes(app.Group(""))
	return router
}
