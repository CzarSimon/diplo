package main

import (
	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

// saveAndBroadcastMessage saves new message and broadcasts it to the channel.
func (env *Env) saveAndBroadcastMessage(msg chat.Message) error {
	err := env.MessageRepo.SaveMessage(msg)
	if err != nil {
		return err
	}
	return env.broadcastMessage(msg)
}

// broadcastMessage sends a message to all the channels active subscribers.
func (env *Env) broadcastMessage(msg chat.Message) error {
	userIDs, err := env.getSubscribedUsers(msg.ChannelID)
	if err != nil {
		return err
	}

	for _, userID := range userIDs {
		msgChannels, ok := env.Channels.GetUserChannels(userID)
		if !ok {
			continue
		}
		sendMessageToChannels(msg, msgChannels)
	}
	return nil
}

// sendMessageToChannels sends a message to a list of channels.
func sendMessageToChannels(msg chat.Message, msgChannels []chan chat.Message) {
	for _, msgChan := range msgChannels {
		msgChan <- msg
	}
}
