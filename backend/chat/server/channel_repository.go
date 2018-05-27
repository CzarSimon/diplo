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
	GetUserChannelsPerGame(userID, gameID string) ([]chat.Channel, error)
	GetSubscribedUsers(channelID string) ([]string, error)
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

const insertChannelQuery = `INSERT INTO channel(id, name, created_at, game_id) VALUES ($1, $2, $3, $4)`

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

const insertChannelMemberQuery = `INSERT INTO channel_member(user_id, channel_id, is_subscribed) VALUES ($1, $2, TRUE)
                                  ON CONFLICT ON CONSTRAINT channel_member_pkey DO
                                  UPDATE SET is_subscribed = TRUE WHERE EXCLUDED.user_id = $1 AND EXCLUDED.channel_id =$2 ;`

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

const deleteChannelMemberQuery = `UPDATE channel_member SET is_subscribed = FALSE WHERE user_id = $1 AND channel_id = $2`

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

const selectUsersChannelsQuery = `SELECT c.id, c.name, c.created_at, c.game_id FROM channel c
                                  INNER JOIN channel_member m ON c.id = m.channel_id
                                  WHERE m.user_id = $1 AND m.is_subscribed = TRUE AND c.game_id = $2`

// GetUserChannelsPerGame gets all channel that a users is connected to in a game.
func (repo *PgChannelRepo) GetUserChannelsPerGame(userID, gameID string) ([]chat.Channel, error) {
	rows, err := repo.db.Query(selectUsersChannelsQuery, userID, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	channels := make([]chat.Channel, 0)
	var channel chat.Channel
	for rows.Next() {
		err = rows.Scan(&channel.ID, &channel.Name, &channel.CreatedAt, &channel.GameID)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}
	return channels, nil
}

const selectChannelMembersQuery = `SELECT user_id FROM channel_member WHERE channel_id = $1 AND is_subscribed = TRUE`

// GetSubscribedUsers gets ids of all users currently subscribed to a channel.
func (repo *PgChannelRepo) GetSubscribedUsers(channelID string) ([]string, error) {
	rows, err := repo.db.Query(selectChannelMembersQuery, channelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	userIDs := make([]string, 0)
	var userID string
	for rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}
