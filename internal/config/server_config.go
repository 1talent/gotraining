package config

import "github.com/1talent/gotraining/internal/util"

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
	Echo EchoServer
}

func DefaultServiceConfigFromEnv() Server {
	return Server{
		Echo: EchoServer{
			ListenAddress: util.GetEnv("SERVER_ECHO_LISTEN_ADDRESS", ":8080"),
		},
	}
}
