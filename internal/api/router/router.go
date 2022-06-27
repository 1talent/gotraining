package router

import (
	"github.com/1talent/gotraining/internal/api"
	"github.com/1talent/gotraining/internal/api/middleware"
	"github.com/labstack/echo/v4"
)

func Init(s *api.Server) {
	s.Echo = echo.New()
	s.Router = &api.Router{
		Routes: nil, // will be populated by handlers.AttachAllRoutes(s)

		// Unsecured base group available at /**
		Root: s.Echo.Group(""),

		// OAuth2, unsecured or secured by bearer auth, available at /api/v1/auth/**
		APIV1Auth: s.Echo.Group("/api/v1/auth", middleware.AuthWithConfig(middleware.AuthConfig{
			S:    s,
			Mode: middleware.AuthModeRequired,
			Skipper: func(c echo.Context) bool {
				switch c.Path() {
				case "/api/v1/auth/forgot-password",
					"/api/v1/auth/forgot-password/complete",
					"/api/v1/auth/login",
					"/api/v1/auth/refresh",
					"/api/v1/auth/register":
					return true
				}
				return false
			},
		})),
	}
}
