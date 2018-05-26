package main

import "log"

// Env holds reference to repositories and configuration.
type Env struct {
	ChannelRepo ChannelRepositoryInterface
	MessageRepo MessageRepositoryInterface
	UserRepo    UserRepositoryInterface
	config      Config
}

// setupEnv sets up environment with postgres repositories.
func setupEnv(config Config) *Env {
	db, err := config.db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return &Env{
		ChannelRepo: NewPgChannelRepo(db),
		MessageRepo: NewPgMessageRepo(db),
		UserRepo:    NewPgUserRepo(db),
		config:      config,
	}
}
