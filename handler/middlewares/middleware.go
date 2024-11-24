package middlewares

import (
	"golang_template/internal/casbin"
	"golang_template/internal/config"
	"log"
)

// Authentication
// JWT

type Middleware interface {
	Auth() AuthMiddleware
	Authorization() AuthorizationMiddleware
}

type middleware struct {
	authMiddleware          AuthMiddleware
	authorizationMiddleware AuthorizationMiddleware
}

func NewMiddleware() Middleware {
	cfg, err := config.LoadConfig("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err.Error())
	}

	casbinService, err := casbin.NewCasbinService()
	if err != nil {
		log.Fatalf("Failed to create casbin service: %v", err.Error())
	}

	authMiddle := NewAuthMiddleware(&cfg.JWT)
	authzMiddle := NewAuthorizationMiddleware(casbinService)
	return &middleware{authMiddleware: authMiddle, authorizationMiddleware: authzMiddle}
}

func (m *middleware) Auth() AuthMiddleware                   { return m.authMiddleware }
func (m *middleware) Authorization() AuthorizationMiddleware { return m.authorizationMiddleware }
