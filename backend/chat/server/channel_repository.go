package main

import (
	"database/sql"

	"github.com/CzarSimon/diplo/backend/chat/pkg/chat"
)

// ChannelRepositoryInterface handles interactions
// with the underlying datastore for channels.
type ChannelRepositoryInterface interface {
	SaveChannel(channel chat.Channel) error
	AddUser(userID, channelID string) error
	RemoveUser(userID, channelID string) error
}

// PgChannelRepo postgres implementation of ChannelRepositoryInterface.
type PgChannelRepo struct {
	db *sql.DB
}

// NewPgChannelRepo creates a new postgres Channel repository.
func NewPgChannelRepo(db *sql.DB) *PgChannelRepo {
	return &PgChannelRepo{
		db: db,
	}
}

const (
	insertChannelQuery       = `INSERT INTO channel(id, name, created_at, game_id) VALUES ($1, $2, $3, $4)`
	insertChannelMemberQuery = `INSERT INTO channel_member(user_id, channel_id, is_subscribed) VALUES ($1, $2, 'TRUE')`
	deleteChannelMemberQuery = `UPDATE channel_member SET is_subscribed = FALSE WHERE user_id = $1 AND channel_id = $2`
)

// SaveChannel creates a new channel and adds the
// specified list of inital users to it.
func (repo *PgChannelRepo) SaveChannel(channel chat.Channel) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		insertChannelQuery, channel.ID, channel.Name, channel.CreatedAt, channel.GameID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = repo.addInitalChannelUsers(channel, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// addInitalChannelUsers Adds inital users to a channel.
func (repo *PgChannelRepo) addInitalChannelUsers(channel chat.Channel, tx *sql.Tx) error {
	stmt, err := tx.Prepare(insertChannelMemberQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, userID := range channel.Recievers {
		_, err = stmt.Exec(userID, channel.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// AddUser records a users as subscirbed to a channel in the database.
func (repo *PgChannelRepo) AddUser(userID, channelID string) error {
	stmt, err := repo.db.Prepare(insertChannelMemberQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, channelID)
	return err
}

// RemoveUser unsubscribes a user from the a channel in the database.
func (repo *PgChannelRepo) RemoveUser(userID, channelID string) error {
	stmt, err := repo.db.Prepare(deleteChannelMemberQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, channelID)
	return err
}
