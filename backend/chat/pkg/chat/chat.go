package chat

import "time"

// Channel represents a many to many message relationship.
type Channel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	GameID    string    `json:"gameId"`
	Recievers []string  `json:"recievers"`
}

// Message content and metadata of a chat message.
type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	ChannelID string    `json:"channelId"`
	AuthorID  string    `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
}
