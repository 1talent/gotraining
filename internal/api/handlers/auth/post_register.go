package auth

import (
	"fmt"

	"github.com/1talent/gotraining/internal/api"
	"github.com/1talent/gotraining/internal/util"
	"github.com/labstack/echo/v4"
)

func PostRegisterRoute(s *api.Server) *echo.Route {
	return s.Router.APIV1Auth.POST("/register", postRegisterHandler(s))
}

func postRegisterHandler(s *api.Server) echo.HandlerFunc {
	return func(c echo.Context) error {

		// boring administrative code that we will need but not really business logic
		ctx := c.Request().Context()
		// log capability with the zerolog library and customisations in log.go and context.go
		log := util.LogFromContext(ctx)
		log.Info().Msg("we have an info logger implemented")

		// this is the critical part - we are parsing out the value from our client
		// now, we come to your original quetsion in the discord channel - how to do "go swagger"
		// we actually want to generate our payload types (payload meaning the payload coming from client)
		// from swagger definitions
		// should be "types" but we use "models" for now until we fix swagger output path
		var body PostRegisterPayLoad
		fmt.Println(body) // TODO: change models to types

		return nil
	}
}
