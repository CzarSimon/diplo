package main

import (
	"database/sql"

	"github.com/CzarSimon/diplo/backend/directory/pkg/directory"
	"github.com/CzarSimon/diplo/backend/pkg/httputil"
	"github.com/lib/pq"
)

// UserRepositoryInterface handles interactions
// with the underlying datastore for users.
type UserRepositoryInterface interface {
	SaveUser(user directory.User) error
	FindUser(userIDorEmail string) (directory.User, error)
	FindUsers(userIDs []string) ([]directory.User, error)
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
  INSERT INTO account(id, email, username, password, given_name, surname, joined_at, ranking)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

// SaveUser save a new user in the database.
func (repo *PgUserRepo) SaveUser(user directory.User) error {
	stmt, err := repo.db.Prepare(insertUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		user.ID, user.Email, user.Username, user.Password,
		user.GivenName, user.Surname, user.JoinedAt, user.Ranking)
	return err
}

const selectUserQuery = `
  SELECT id, email, username, password, given_name, surname, joined_at, ranking
  FROM account WHERE id = $1 OR email = $1`

// FindUser checks for and returns a user id if present in the database.
func (repo *PgUserRepo) FindUser(userIDorEmail string) (directory.User, error) {
	var user directory.User
	err := repo.db.QueryRow(selectUserQuery, userIDorEmail).Scan(
		&user.ID, &user.Email, &user.Username, &user.Password,
		&user.GivenName, &user.Surname, &user.JoinedAt, &user.Ranking)
	if err == sql.ErrNoRows {
		return user, httputil.ErrUserNotFound
	}
	return user, err
}

const selectUsersQuery = `
  SELECT id, email, username, password, given_name, surname, joined_at, ranking
  FROM account WHERE id = ANY($1)`

// FindUsers returns users matching the supplied list of user ids.
func (repo *PgUserRepo) FindUsers(userIDs []string) ([]directory.User, error) {
	rows, err := repo.db.Query(selectUsersQuery, pq.Array(userIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return createUsersFromRows(rows)
}

// createUsersFromRows creates a slice of users from a query.
func createUsersFromRows(rows *sql.Rows) ([]directory.User, error) {
	users := make([]directory.User, 0)
	var user directory.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password,
			&user.GivenName, &user.Surname, &user.JoinedAt, &user.Ranking)
		if err != nil {
			return nil, err
		}
		users = append(users, sanitizeUser(user))
	}
	return users, nil
}
