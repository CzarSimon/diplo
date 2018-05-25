package main

import (
	"database/sql"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

// MessageRepositoryInterface handles interactions with the underlying
// datastore for messages.
type MessageRepositoryInterface interface {
	SaveMessage(msg chat.Message) error
}

// PgMessageRepo postgres implementation of MessageRepositoryInterface.
type PgMessageRepo struct {
	db *sql.DB
}

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
