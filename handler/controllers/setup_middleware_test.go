package controllers

import "github.com/gofiber/fiber/v2"

type MockMiddleware struct {
	mockAuthMiddleware MockAuthMiddleware
}

func (m *MockMiddleware) Auth() MockAuthMiddleware { return m.mockAuthMiddleware }

type MockAuthMiddleware struct {
	forceLoggedIn bool
}

func (ma *MockAuthMiddleware) Handle() fiber.Handler {
	if ma.forceLoggedIn {
		return func(ctx *fiber.Ctx) error {
			return ctx.Next()
		}
	} else {
		return nil
	}
}
