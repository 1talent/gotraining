package util

import "context"

type contextKey string

const (
	CTXKeyDisableLogger contextKey = "disable_logger"
)

func ShouldDisableLogger(ctx context.Context) bool {
	s := ctx.Value(CTXKeyDisableLogger)
	if s == nil {
		return false
	}

	shouldDisable, ok := s.(bool)
	if !ok {
		return false
	}
	return shouldDisable
}
