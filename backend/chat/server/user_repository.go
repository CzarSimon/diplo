package main

import (
	"database/sql"
)

// UserRepositoryInterface handles interactions
// with the underlying datastore for users.
type UserRepositoryInterface interface {
	SaveUser(userID string) error
	RemoveUser(userID string) error
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

const (
	insertUserQuery = `INSERT INTO chat_user(id) VALUES ($1)`
	deleteUserQuery = `DELETE FROM chat_user WHERE id = $1`
)

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
