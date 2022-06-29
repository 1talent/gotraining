package router

import (
	"github.com/1talent/gotraining/internal/api"
	"github.com/1talent/gotraining/internal/api/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func Init(s *api.Server) {
	s.Echo = echo.New()
	s.Echo.HideBanner = true
	s.Echo.HTTPErrorHandler = HTTPErrorHandlerWithConfig(true)

	// General middleware
	if s.Config.Echo.EnableTrailingSlashMiddleware {
		s.Echo.Pre(echoMiddleware.RemoveTrailingSlash())
	} else {
		log.Warn().Msg("Disabling trailing slash middleware due to environment config")
	}

	if s.Config.Echo.EnableRecoverMiddleware {
		s.Echo.Use(echoMiddleware.Recover())
	} else {
		log.Warn().Msg("Disabling recover middleware due to environment config")
	}

	if s.Config.Echo.EnableSecureMiddleware {
		s.Echo.Use(echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
			Skipper:               echoMiddleware.DefaultSecureConfig.Skipper,
			XSSProtection:         s.Config.Echo.SecureMiddleware.XSSProtection,
			ContentTypeNosniff:    s.Config.Echo.SecureMiddleware.ContentTypeNosniff,
			XFrameOptions:         s.Config.Echo.SecureMiddleware.XFrameOptions,
			HSTSMaxAge:            s.Config.Echo.SecureMiddleware.HSTSMaxAge,
			HSTSExcludeSubdomains: s.Config.Echo.SecureMiddleware.HSTSExcludeSubdomains,
			ContentSecurityPolicy: s.Config.Echo.SecureMiddleware.ContentSecurityPolicy,
			CSPReportOnly:         s.Config.Echo.SecureMiddleware.CSPReportOnly,
			HSTSPreloadEnabled:    s.Config.Echo.SecureMiddleware.HSTSPreloadEnabled,
			ReferrerPolicy:        s.Config.Echo.SecureMiddleware.ReferrerPolicy,
		}))
	} else {
		log.Warn().Msg("Disabling secure middleware due to environment config")
	}

	if s.Config.Echo.EnableRequestIDMiddleware {
		s.Echo.Use(echoMiddleware.RequestID())
	} else {
		log.Warn().Msg("Disabling request ID middleware due to environment config")
	}

	if s.Config.Echo.EnableCORSMiddleware {
		s.Echo.Use(echoMiddleware.CORS())
	} else {
		log.Warn().Msg("Disabling CORS middleware due to environment config")
	}

	// TODO: Implement specific middlewares: Logger, CacheControl, Profiler, Auth

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
