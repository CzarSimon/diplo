package main

import (
	"database/sql"
	"time"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

// MessageRepositoryInterface handles interactions with the underlying
// datastore for messages.
type MessageRepositoryInterface interface {
	SaveMessage(msg chat.Message) error
	GetMessagesSince(channelID string, since time.Time) ([]chat.Message, error)
}

// PgMessageRepo postgres implementation of MessageRepositoryInterface.
type PgMessageRepo struct {
	db *sql.DB
}

// NewPgMessageRepo creates a new postgres Message repository.
func NewPgMessageRepo(db *sql.DB) *PgMessageRepo {
	return &PgMessageRepo{
		db: db,
	}
}

const insertMessageQuery = `
  INSERT INTO message(id, message_text, channel_id, author, created_at) VALUES ($1, $2, $3, $4, $5)`

// SaveMessage saves a message in the postgres database.
func (repo *PgMessageRepo) SaveMessage(msg chat.Message) error {
	stmt, err := repo.db.Prepare(insertMessageQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(msg.ID, msg.Text, msg.ChannelID, msg.AuthorID, msg.CreatedAt)
	return err
}

const getMessagesSinceQuery = `
  SELECT id, message_text, channel_id, author, created_at FROM message WHERE channel_id = $1 AND created_at >= $2 ORDER BY created_at`

// GetMessagesSince get messages in a channel since a certain date.
func (repo *PgMessageRepo) GetMessagesSince(channelID string, since time.Time) ([]chat.Message, error) {
	rows, err := repo.db.Query(getMessagesSinceQuery, channelID, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return createMessagesFromRows(rows)
}

// createMessagesFromRows organizes sql result set into a list of messages.
func createMessagesFromRows(rows *sql.Rows) ([]chat.Message, error) {
	messages := make([]chat.Message, 0)
	var msg chat.Message
	for rows.Next() {
		err := rows.Scan(&msg.ID, &msg.Text, &msg.ChannelID, &msg.AuthorID, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
