package main

import "log"

// Env holds reference to repositories and configuration.
type Env struct {
	MessageRepo MessageRepositoryInterface
	config      Config
}

// setupEnv sets up environment with postgres repositories.
func setupEnv(config Config) *Env {
	db, err := config.db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return &Env{
		MessageRepo: NewPgMessageRepo(db),
		config:      config,
	}
}
