package main

import (
	"log"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

// saveAndBroadcastMessage saves new message and broadcasts it to the channel
func saveAndBroadcastMessage(message chat.Message) error {
	log.Printf("Broadcasting: %s\n", message)
	return nil
}
