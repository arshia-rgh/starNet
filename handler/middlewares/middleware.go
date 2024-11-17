package middlewares

import (
	"golang_template/internal/config"
	"log"
)

// Authentication
// JWT

type Middleware interface {
	Auth() AuthMiddleware
}

type middleware struct {
	authMiddleware AuthMiddleware
}

func NewMiddleware() Middleware {
	cfg, err := config.LoadConfig("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err.Error())
	}
	authMiddle := NewAuthMiddleware(&cfg.JWT)
	return &middleware{authMiddleware: authMiddle}
}

func (m *middleware) Auth() AuthMiddleware { return m.authMiddleware }
