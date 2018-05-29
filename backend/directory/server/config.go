package main

import (
	"crypto/sha256"
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
	return Config{
		jwtSecret: []byte(jwtSecret),
		saltKey:   getSaltKey(),
		db:        endpoint.NewPGConfig(DB_NAME),
		server:    endpoint.NewServerAddr(SERVER_NAME),
	}
}

func getSaltKey() []byte {
	saltKey := os.Getenv(SALT_KEY_KEY)
	if saltKey == "" {
		log.Fatalf("%s not provided", SALT_KEY_KEY)
	}
	return sha256.New().Sum([]byte(saltKey))[0:32]
}
