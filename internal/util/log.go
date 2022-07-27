package util

import (
	"context"

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
