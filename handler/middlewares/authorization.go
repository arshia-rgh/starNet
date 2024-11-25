package middlewares

import (
	"golang_template/internal/casbin"

	"github.com/gofiber/fiber/v2"
)

type AuthorizationMiddleware interface {
	Handle() fiber.Handler
}

type authorizationMiddleware struct {
	casbinService casbin.CasbinService
}

func NewAuthorizationMiddleware(casbinService casbin.CasbinService) AuthorizationMiddleware {
	return &authorizationMiddleware{
		casbinService: casbinService,
	}
}

func (a *authorizationMiddleware) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("userRole").(string)
		path := c.Path()
		method := c.Method()

		allowed, err := a.casbinService.Enforce(userRole, path, method)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Authorization error",
			})
		}

		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Access denied",
			})
		}

		return c.Next()
	}
}
