package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang_template/internal/config"
	"golang_template/pkg"
	"strings"
)

type AuthMiddleware interface {
	Handle() fiber.Handler
}

type authMiddleware struct {
	cfg *config.JWTConfig
}

func NewAuthMiddleware(cfg *config.JWTConfig) AuthMiddleware {
	return &authMiddleware{cfg: cfg}
}

func (a *authMiddleware) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "credentials not provided"})
		}
		token := strings.TrimPrefix(authorization, "Bearer")

		userID, err := pkg.VerifyToken(token, a.cfg.Secret)
		if err != nil {
			if errors.Is(err, pkg.ErrInvalidToken) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "wrong token",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed validating token"})
		}
		c.Locals("user", userID)
		return c.Next()

	}
}
