package directory

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/CzarSimon/diplo/backend/pkg/id"
)

const (
	StartingRank = 1000.0
	SaltLength   = 100
)

// User datatype for representing a user.
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
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
		Salt:      generateSalt(),
		GivenName: user.GivenName,
		Surname:   user.Surname,
		JoinedAt:  time.Now().UTC(),
		Ranking:   StartingRank,
	}
}

// generateSalt craetes a random salt to be userd to hash a a users password.
func generateSalt() string {
	bytes := make([]byte, SaltLength)
	_, err := rand.Read(bytes)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal()
	}
	return base64.StdEncoding.EncodeToString(bytes)
}
