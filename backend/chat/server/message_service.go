package main

import (
	"log"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

func saveAndBroadcastNewMessage(text, channelID, userID string) (chat.Message, error) {
	message := chat.NewMessage(text, channelID, userID)
	return message, saveAndBroadcastMessage(message)
}

func saveAndBroadcastMessage(message chat.Message) error {
	log.Printf("Broadcasting: %s\n", message)
	return nil
}
