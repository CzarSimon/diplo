package main

import (
	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

// saveAndBroadcastMessage saves new message and broadcasts it to the channel
func (env *Env) saveAndBroadcastMessage(msg chat.Message) error {
	return env.MessageRepo.SaveMessage(msg)
}
