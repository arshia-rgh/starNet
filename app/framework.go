package app

import "github.com/gofiber/fiber/v2"

func (a *application) InitFramework() *fiber.App {
	return fiber.New(fiber.Config{BodyLimit: 100 * 1024 * 1024})
}
