package handlers

import (
	"github.com/1talent/gotraining/internal/api"
	"github.com/1talent/gotraining/internal/api/handlers/auth"
	"github.com/labstack/echo/v4"
)

func AttachAllRoutes(s *api.Server) {
	s.Router.Routes = []*echo.Route{
		auth.PostRegisterRoute(s),
		// we will keep adding
	}
}
