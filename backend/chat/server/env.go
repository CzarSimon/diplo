package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
)

// Env holds reference to repositories and configuration.
type Env struct {
	ChannelRepo ChannelRepositoryInterface
	MessageRepo MessageRepositoryInterface
	UserRepo    UserRepositoryInterface
	AuthOpts    *httputil.AuthOptions
	Channels    *UserChannels
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
		AuthOpts:    httputil.NewAuthOptions(config.jwtSecret, "directory", config.authExemptedRoutes...),
		Channels:    NewUserChannels(),
		config:      config,
	}
}

// Common session errors
var ErrDuplicateSesion = httputil.NewError(http.StatusInternalServerError, "Session already exists")

// SessionChannels is a map of session ids and message channels.
type SessionChannels map[string]chan chat.Message

// Add adds a message channel identifed by a session id.
func (sc SessionChannels) Add(sessionID string, msgChan chan chat.Message) error {
	_, ok := sc[sessionID]
	if ok {
		return ErrDuplicateSesion
	}
	sc[sessionID] = msgChan
	return nil
}

// UserChannels map of user ids to channels for each of their sessions
type UserChannels struct {
	lock     sync.RWMutex
	channels map[string]SessionChannels
}

func NewUserChannels() *UserChannels {
	return &UserChannels{
		lock:     sync.RWMutex{},
		channels: make(map[string]SessionChannels),
	}
}

// Add adds a message channel for a user session.
func (uc *UserChannels) Add(userID, sessionID string, msgChan chan chat.Message) error {
	uc.lock.Lock()
	defer uc.lock.Unlock()
	sessionChannels, ok := uc.channels[userID]
	if !ok {
		sessionChannels = make(SessionChannels)
	}

	err := sessionChannels.Add(sessionID, msgChan)
	if err != nil {
		return err
	}
	uc.channels[userID] = sessionChannels
	return nil
}

// Delete
func (uc *UserChannels) Delete(userID, sessionID string) {
	uc.lock.Lock()
	defer uc.lock.Unlock()
	sessionChannels, ok := uc.channels[userID]
	if !ok {
		return
	}

	delete(sessionChannels, sessionID)
	if len(sessionChannels) < 1 {
		delete(uc.channels, userID)
		return
	}
	uc.channels[userID] = sessionChannels
}

// GetUserChannels gets all channels registered to a user.
func (uc *UserChannels) GetUserChannels(userID string) ([]chan chat.Message, bool) {
	uc.lock.RLock()
	defer uc.lock.RUnlock()
	sessionChannels, ok := uc.channels[userID]
	if !ok {
		return nil, false
	}

	msgChannels := make([]chan chat.Message, 0, len(sessionChannels))
	for _, msgChannel := range sessionChannels {
		msgChannels = append(msgChannels, msgChannel)
	}
	return msgChannels, true
}
