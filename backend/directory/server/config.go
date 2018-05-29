package main

import (
	"log"
	"os"

	endpoint "github.com/CzarSimon/go-endpoint"
)

const (
	SERVER_NAME    = "DIPLO_DIRECTORY_SERVER"
	DB_NAME        = "DIPLO_DIRECTORY_DB"
	JWT_SECRET_KEY = "JWT_SECRET"
	SALT_KEY_KEY   = "SALT_KEY"
)

// Config holds configuration values.
type Config struct {
	jwtSecret []byte
	saltKey   []byte
	server    endpoint.ServerAddr
	db        endpoint.SQLConfig
}

// GetConfig gets configuration values from the environment.
func GetConfig() Config {
	jwtSecret := os.Getenv(JWT_SECRET_KEY)
	if jwtSecret == "" {
		log.Fatalf("%s not provided", JWT_SECRET_KEY)
	}
	saltKey := os.Getenv(SALT_KEY_KEY)
	if saltKey == "" {
		log.Fatalf("%s not provided", SALT_KEY_KEY)
	}
	return Config{
		jwtSecret: []byte(jwtSecret),
		saltKey:   []byte(saltKey),
		db:        endpoint.NewPGConfig(DB_NAME),
		server:    endpoint.NewServerAddr(SERVER_NAME),
	}
}
