package main

import endpoint "github.com/CzarSimon/go-endpoint"

const (
	SERVER_NAME = "DIPLO_CHAT_SERVER"
	DB_NAME     = "DIPLO_CHAT_DB"
)

// Config holds configuration values.
type Config struct {
	server endpoint.ServerAddr
	db     endpoint.SQLConfig
}

// GetConfig gets configuration values from the environment.
func GetConfig() Config {
	return Config{
		db:     endpoint.NewPGConfig(DB_NAME),
		server: endpoint.NewServerAddr(SERVER_NAME),
	}
}
