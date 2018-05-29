package main

import (
	"database/sql"

	"github.com/CzarSimon/diplo/backend/directory/pkg/directory"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
)

// UserRepositoryInterface handles interactions
// with the underlying datastore for users.
type UserRepositoryInterface interface {
	SaveUser(user directory.User) error
	FindUser(userID string) (directory.User, error)
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

const insertUserQuery = `
  INSERT INTO app_user(id, email, username, password, salt, given_naem, surname, joined_at, ranking)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

// SaveUser save a new user in the database.
func (repo *PgUserRepo) SaveUser(user directory.User) error {
	stmt, err := repo.db.Prepare(insertUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	return err
}

const selectUserQuery = "SELECT id FROM app_user WHERE id = $1"

// FindUser checks for and returns a user id if present in the database.
func (repo *PgUserRepo) FindUser(userID string) (directory.User, error) {
	var user directory.User
	err := repo.db.QueryRow(selectUserQuery, userID).Scan(&user.ID)
	if err == sql.ErrNoRows {
		return user, httputil.ErrUserNotFound
	}
	return user, err
}
