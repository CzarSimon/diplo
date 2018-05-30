package main

import (
	"time"

	"github.com/CzarSimon/diplo/backend/directory/pkg/directory"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/square/go-jose.v2/jwt"
)

// saveUser creates new user, encrypts sensitive information and stores it.
func (env *Env) saveUser(user directory.User) (directory.User, error) {
	newUser := directory.NewUser(user)
	hashedPassword, err := hashPassword(newUser)
	if err != nil {
		return user, err
	}
	newUser.Password = hashedPassword
	err = env.UserRepo.SaveUser(newUser)
	if err != nil {
		return user, err
	}
	return sanitizeUser(newUser), nil
}

// loginUser checks the login credentials of a user and issues a JWT if correct.
func (env *Env) loginUser(user directory.User) (directory.Token, error) {
	storedUser, err := env.getUserByEmail(user.Email)
	if err != nil {
		return directory.EmptyToken, err
	}

	err = env.checkUserPassword(user.Password, storedUser)
	if err != nil {
		return directory.EmptyToken, err
	}

	return env.createJWTToken(storedUser)
}

// getUserByEmail retrieves a user by email.
func (env *Env) getUserByEmail(email string) (directory.User, error) {
	user, err := env.UserRepo.FindUser(email)
	if err != nil {
		return directory.EmptyUser, err
	}
	return user, nil
}

// checkUserPassword retrieves user from database and verifies the provided password.
func (env *Env) checkUserPassword(candidatePwd string, user directory.User) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(candidatePwd))
	if err != nil {
		return ErrAccessDenied
	}
	return nil
}

func (env *Env) createJWTToken(user directory.User) (directory.Token, error) {
	claims := jwt.Claims{
		Subject:   user.ID,
		Issuer:    "directory",
		NotBefore: jwt.NewNumericDate(time.Now().UTC().Add(-1 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		Expiry:    jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
	}

	tokenString, err := jwt.Signed(env.config.signer).Claims(claims).CompactSerialize()
	if err != nil {
		return directory.EmptyToken, err
	}
	return directory.NewToken(tokenString), nil
}

// sanitizeUser removes sensitive information from a user.
func sanitizeUser(user directory.User) directory.User {
	user.Password = ""
	user.Salt = ""
	return user
}

// hashPassword hashes a users password with an encrypted salt.
func hashPassword(user directory.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
