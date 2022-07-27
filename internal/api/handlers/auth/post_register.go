package auth

import (
	"github.com/1talent/gotraining/internal/api"
	"github.com/1talent/gotraining/internal/util"
	"github.com/labstack/echo/v4"
)

func PostRegisterRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Auth.POST("/register", postRegisterHandler(s))
}

func postRegisterHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		// log capability with the zerolog library and customisations in log.go and context.go
		log := util.LogFromContext(ctx)
		log.Info().Msg("we have an info logger implemented")
		return nil
	}
}
