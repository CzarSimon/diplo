package chat

import (
	"fmt"
	"time"

	"github.com/CzarSimon/diplo/backend/pkg/id"
)

// Channel represents a many to many message relationship.
type Channel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	GameID    string    `json:"gameId"`
	Recievers []string  `json:"recievers"`
}

// NewChannel creates a new channel.
func NewChannel(name, gameID string) Channel {
	return Channel{
		ID:        id.New(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		GameID:    gameID,
	}
}

// String returns a string representation of a channel.
func (c *Channel) String() string {
	return fmt.Sprintf("Channel: %s\nMetadata: id=%s createdAt=%v gameId=%s",
		c.Name, c.ID, c.CreatedAt, c.GameID)
}

// Message content and metadata of a chat message.
type Message struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	ChannelID string    `json:"channelId"`
	AuthorID  string    `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewMessage creates a new message.
func NewMessage(text, channelID, authorID string) Message {
	return Message{
		ID:        id.New(),
		Text:      text,
		ChannelID: channelID,
		AuthorID:  authorID,
		CreatedAt: time.Now().UTC(),
	}
}

// String returns a string representation of a message.
func (m *Message) String() string {
	return fmt.Sprintf("Message: %s\nMetadata: id=%s, channel=%s, author=%s, createdAt=%v",
		m.Text, m.ID, m.ChannelID, m.AuthorID, m.CreatedAt)
}
