package config

import "net/http"

type ServerConfig struct {
	Address string       // 服务器地址
	Handler http.Handler // 服务器 Handler
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Address: "localhost:8080",
		Handler: nil,
	}
}
