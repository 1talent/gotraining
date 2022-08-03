package util

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogFromContext(ctx context.Context) *zerolog.Logger {
	l := log.Ctx(ctx)
	if l.GetLevel() == zerolog.Disabled {
		if ShouldDisableLogger(ctx) {
			return l
		}
		l = &log.Logger
	}
	return l
}

// LogFromEchoContext returns a request-specific zerolog instance using the echo.Context of the request.
// The returned logger will have the request ID as well as some other value predefined.
func LogFromEchoContext(c echo.Context) *zerolog.Logger {
	return LogFromContext(c.Request().Context())
}
