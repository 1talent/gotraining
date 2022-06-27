package config

import "github.com/1talent/gotraining/internal/util"

type EchoServer struct {
	ListenAddress string
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
