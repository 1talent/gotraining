package api

import (
	"context"
	"database/sql"

	"github.com/1talent/gotraining/internal/config"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
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
	DB     *sql.DB
}

func NewServer(config config.Server) *Server {
	return &Server{
		Config: config,
		DB:     nil,
	}
}

func (s *Server) Start() error {
	return s.Echo.Start(s.Config.Echo.ListenAddress)
}

func (s *Server) InitDB(ctx context.Context) error {
	db, err := sql.Open("postgres", s.Config.Database.ConnectionString())
	if err != nil {
		return err
	}

	if s.Config.Database.MaxOpenConns > 0 {
		db.SetMaxOpenConns(s.Config.Database.MaxOpenConns)
	}
	if s.Config.Database.MaxIdleConns > 0 {
		db.SetMaxIdleConns(s.Config.Database.MaxIdleConns)
	}
	if s.Config.Database.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(s.Config.Database.ConnMaxLifetime)
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	s.DB = db

	return nil
}
