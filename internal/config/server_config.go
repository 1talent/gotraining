package config

import (
	"runtime"
	"time"

	"github.com/1talent/gotraining/internal/util"
)

type EchoServer struct {
	ListenAddress                 string
	EnableCORSMiddleware          bool
	EnableLoggerMiddleware        bool
	EnableRecoverMiddleware       bool
	EnableRequestIDMiddleware     bool
	EnableTrailingSlashMiddleware bool
	EnableSecureMiddleware        bool
	EnableCacheControlMiddleware  bool
	SecureMiddleware              EchoServerSecureMiddleware
}

// EchoServerSecureMiddleware represents a subset of echo's secure middleware config relevant to the app server.
// https://github.com/labstack/echo/blob/master/middleware/secure.go
type EchoServerSecureMiddleware struct {
	XSSProtection         string
	ContentTypeNosniff    string
	XFrameOptions         string
	HSTSMaxAge            int
	HSTSExcludeSubdomains bool
	ContentSecurityPolicy string
	CSPReportOnly         bool
	HSTSPreloadEnabled    bool
	ReferrerPolicy        string
}

type Server struct {
	Echo     EchoServer
	Database Database
}

func DefaultServiceConfigFromEnv() Server {
	return Server{
		Database: Database{
			Host:     util.GetEnv("PGHOST", "localhost"),
			Port:     util.GetEnvAsInt("PGPORT", 5433),
			Database: util.GetEnv("PGDATABASE", "development"),
			Username: util.GetEnv("PGUSER", "dbuser"),
			Password: util.GetEnv("PGPASSWORD", "dbpass"),
			AdditionalParams: map[string]string{
				"sslmode": util.GetEnv("PGSSLMODE", "disable"),
			},
			MaxOpenConns:    util.GetEnvAsInt("DB_MAX_OPEN_CONNS", runtime.NumCPU()*2),
			MaxIdleConns:    util.GetEnvAsInt("DB_MAX_IDLE_CONNS", 1),
			ConnMaxLifetime: time.Second * time.Duration(util.GetEnvAsInt("DB_CONN_MAX_LIFETIME_SEC", 60)),
		},
		Echo: EchoServer{
			ListenAddress: util.GetEnv("SERVER_ECHO_LISTEN_ADDRESS", ":8050"),
		},
	}
}
