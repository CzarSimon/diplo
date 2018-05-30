package main

import (
	"log"
	"os"

	endpoint "github.com/CzarSimon/go-endpoint"
	jose "gopkg.in/square/go-jose.v2"
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
	signer    jose.Signer
	server    endpoint.ServerAddr
	db        endpoint.SQLConfig
}

// GetConfig gets configuration values from the environment.
func GetConfig() Config {
	jwtSecret, signer := getTokenSignerAndSecret()
	return Config{
		jwtSecret: []byte(jwtSecret),
		signer:    signer,
		db:        endpoint.NewPGConfig(DB_NAME),
		server:    endpoint.NewServerAddr(SERVER_NAME),
	}
}

func getTokenSignerAndSecret() ([]byte, jose.Signer) {
	jwtSecret := os.Getenv(JWT_SECRET_KEY)
	if jwtSecret == "" {
		log.Fatalf("%s not provided", JWT_SECRET_KEY)
	}
	secretBytes := []byte(jwtSecret)
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: secretBytes}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return secretBytes, signer
}
