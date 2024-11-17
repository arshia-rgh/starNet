package middlewares

import "golang_template/internal/config"

// Authentication
// JWT

type Middleware interface {
	Auth() AuthMiddleware
}

type middleware struct {
	authMiddleware AuthMiddleware
}

func NewMiddleware(cfg *config.JWTConfig) Middleware {
	authMiddle := NewAuthMiddleware(cfg)
	return &middleware{authMiddleware: authMiddle}
}

func (m *middleware) Auth() AuthMiddleware { return m.authMiddleware }
