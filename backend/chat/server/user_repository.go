package main

import (
	"database/sql"

	"github.com/CzarSimon/diplo/backend/pkg/httputil"
)

// UserRepositoryInterface handles interactions
// with the underlying datastore for users.
type UserRepositoryInterface interface {
	SaveUser(userID string) error
	RemoveUser(userID string) error
	FindUser(userID string) (string, error)
}

// PgUserRepo postgres implementation of UserRepositoryInterface.
type PgUserRepo struct {
	db *sql.DB
}

// NewPgUserRepo creates a new postgres User repository.
func NewPgUserRepo(db *sql.DB) *PgUserRepo {
	return &PgUserRepo{
		db: db,
	}
}

const insertUserQuery = "INSERT INTO chat_user(id) VALUES ($1)"

// SaveUser save a new user in the database.
func (repo *PgUserRepo) SaveUser(userID string) error {
	stmt, err := repo.db.Prepare(insertUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID)
	return err
}

const deleteUserQuery = "DELETE FROM chat_user WHERE id = $1"

// RemoveUser deletes a user from the database.
func (repo *PgUserRepo) RemoveUser(userID string) error {
	stmt, err := repo.db.Prepare(deleteUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID)
	return err
}

const selectUserQuery = "SELECT id FROM chat_user WHERE id = $1"

// FindUser checks for and returns a user id if present in the database.
func (repo *PgUserRepo) FindUser(userID string) (string, error) {
	var foundID string
	err := repo.db.QueryRow(selectUserQuery, userID).Scan(&foundID)
	if err == sql.ErrNoRows {
		return "", httputil.ErrUserNotFound
	}
	return foundID, err
}
