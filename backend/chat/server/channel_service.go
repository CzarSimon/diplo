package main

import "github.com/CzarSimon/diplo/backend/chat/pkg/chat"

// saveAndRegisterChannel creates a new channel and
// registers it for message broadcasts.
func (env *Env) saveAndRegisterChannel(channel chat.Channel) error {
	return env.ChannelRepo.SaveChannel(channel)
}

// addUserToChannel adds a user to a channel.
func (env *Env) addUserToChannel(userID, channelID string) error {
	return env.ChannelRepo.AddUser(userID, channelID)
}

// removeUserFromChannel adds a user to a channel.
func (env *Env) removeUserFromChannel(userID, channelID string) error {
	return env.ChannelRepo.RemoveUser(userID, channelID)
}

// getSubscribedUsers gets the ids of all users subscribed to a channel.
func (env *Env) getSubscribedUsers(channelID string) ([]string, error) {
	return env.ChannelRepo.GetSubscribedUsers(channelID)
}

// getUsersGameChannels get all channels in a game that a user is subscribed to.
func (env *Env) getUsersGameChannels(userID, gameID string) ([]chat.Channel, error) {
	return env.ChannelRepo.GetUserChannelsPerGame(userID, gameID)
}
