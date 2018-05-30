package directory

import (
	"time"

	"github.com/CzarSimon/diplo/backend/pkg/id"
)

const (
	StartingRank = 1000.0
)

var (
	EmptyUser  = User{}
	EmptyToken = Token{}
)

// User datatype for representing a user.
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	GivenName string    `json:"givenName"`
	Surname   string    `json:"surname"`
	JoinedAt  time.Time `json:"joinedAt"`
	Ranking   float32   `json:"ranking"`
}

// NewUser creates a new user.
func NewUser(user User) User {
	return User{
		ID:        id.New(),
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		GivenName: user.GivenName,
		Surname:   user.Surname,
		JoinedAt:  time.Now().UTC(),
		Ranking:   StartingRank,
	}
}

// Token holds JWT token data.
type Token struct {
	Token string `json:"token"`
}

// NewToken creates a new token.
func NewToken(token string) Token {
	return Token{
		Token: token,
	}
}
