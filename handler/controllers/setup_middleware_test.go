package controllers

import "github.com/gofiber/fiber/v2"

type MockMiddleware struct {
	mockAuthMiddleware MockAuthMiddleware
}

func (m *MockMiddleware) Auth() MockAuthMiddleware { return m.mockAuthMiddleware }

type MockAuthMiddleware struct {
	forceLoggedIn bool
	role          string
}

func (ma *MockAuthMiddleware) Handle() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ma.forceLoggedIn {
			ctx.Locals("userRole", ma.role)
			return ctx.Next()
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}
}
