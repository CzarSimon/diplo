package main

import (
	"database/sql"
	"log"
)

// Env holds reference to repositories and configuration.
type Env struct {
	UserRepo UserRepositoryInterface
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
		config:   config,
		db:       db,
	}
}
