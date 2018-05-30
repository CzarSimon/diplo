package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	jose "gopkg.in/square/go-jose.v2"
)

// Env holds reference to repositories and configuration.
type Env struct {
	UserRepo UserRepositoryInterface
	signer   jose.Signer
	authOpts *httputil.AuthOptions
	config   Config
	db       *sql.DB
}

// setupEnv sets up environment with postgres repositories.
func setupEnv(config Config) *Env {
	db, err := config.db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return &Env{
		UserRepo: NewPgUserRepo(db),
		signer:   setupTokenSigner(config.jwtSecret),
		authOpts: setupAuthOptions(string(config.jwtSecret), "directory"),
		config:   config,
		db:       db,
	}
}

// setupTokenSigner sets up a signer for JWT tokens.
func setupTokenSigner(jwtSecret []byte) jose.Signer {
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: jwtSecret}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return signer
}

// setupAuthOptions sets up options for authentication.
func setupAuthOptions(jwtSecret, tokenIssuer string) *httputil.AuthOptions {
	opts := httputil.NewAuthOptions(jwtSecret, tokenIssuer)
	opts.ValidationLeeway = 1 * time.Hour
	return opts
}
