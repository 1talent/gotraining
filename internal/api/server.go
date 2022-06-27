package api

import (
	"github.com/1talent/gotraining/internal/config"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Routes    []*echo.Route
	Root      *echo.Group
	APIV1Auth *echo.Group
}

type Server struct {
	Config config.Server
	Echo   *echo.Echo
	Router *Router
}

func NewServer(config config.Server) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Start() error {
	return s.Echo.Start(s.Config.Echo.ListenAddress)
}
