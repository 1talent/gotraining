package middleware

import (
	"github.com/1talent/gotraining/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// AuthMode controls the type of authentication check performed for a specific route or group
type AuthMode int

const (
	// AuthModeRequired requires an auth token to be present and valid in order to access the route or group
	AuthModeRequired AuthMode = iota
	// AuthModeSecure requires an auth token to be present and for the user to have recently re-confirmed their authentication in order to access the route or group
	AuthModeSecure
	// AuthModeOptional does not require an auth token to be present, however if it is, it must be valid in order to access the route or group
	AuthModeOptional
	// AuthModeTry does not require an auth token to be present in order to access the route or group and will process the request even if an invalid one has been provided
	AuthModeTry
	// AuthModeNone does not require an auth token to be present in order to access the route or group and will not attempt to parse any authentication provided
	AuthModeNone
)

type AuthConfig struct {
	S    *api.Server // API server used for database and service access
	Mode AuthMode    // Controls type of authentication required (default: AuthModeRequired)
	// FailureMode     AuthFailureMode          // Controls response on auth failure (default: AuthFailureModeUnauthorized)
	// TokenSource     AuthTokenSource          // Sets source of auth token (default: AuthTokenSourceHeader)
	// TokenSourceKey  string                   // Sets key for auth token source lookup (default: "Authorization")
	// Scheme          string                   // Sets required token scheme (default: "Bearer")
	Skipper middleware.Skipper // Controls skipping of certain routes (default: no skipped routes)
	// FormatValidator AuthTokenFormatValidator // Validates the format of the token retrieved
	// TokenValidator  AuthTokenValidator       // Validates token retrieved and returns associated user (default: performs lookup in access_tokens table)
	// Scopes          []string                 // List of scopes required to access endpoint (default: none required)
}

var (
	DefaultAuthConfig = AuthConfig{
		Mode: AuthModeRequired,
		// FailureMode:     AuthFailureModeUnauthorized,
		// TokenSource:     AuthTokenSourceHeader,
		// TokenSourceKey:  echo.HeaderAuthorization,
		// Scheme:          "Bearer",
		Skipper: middleware.DefaultSkipper,
		// FormatValidator: DefaultAuthTokenFormatValidator,
		// TokenValidator:  DefaultAuthTokenValidator,
		// Scopes:          []string{auth.AuthScopeApp.String()},
	}
)

func AuthWithConfig(config AuthConfig) echo.MiddlewareFunc {
	if config.S == nil {
		panic("auth middleware: server is required")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
