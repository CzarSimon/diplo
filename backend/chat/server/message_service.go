package main

import (
	"log"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

// saveAndBroadcastMessage saves new message and broadcasts it to the channel
func (env *Env) saveAndBroadcastMessage(msg chat.Message) error {
	log.Printf("Broadcasting: %s\n", msg)
	return env.MessageRepo.SaveMessage(msg)
}
