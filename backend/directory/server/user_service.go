package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/CzarSimon/diplo/backend/directory/pkg/directory"
	"golang.org/x/crypto/bcrypt"
)

// saveUser creates new user, encrypts sensitive information and stores it.
func (env *Env) saveUser(user directory.User) (directory.User, error) {
	newUser := directory.NewUser(user)
	hashedPassword, err := hashPassword(newUser, env.config.saltKey)
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

// sanitizeUser removes sensitive information from a user.
func sanitizeUser(user directory.User) directory.User {
	user.Password = ""
	user.Salt = ""
	return user
}

// hashPassword hashes a users password with an encrypted salt.
func hashPassword(user directory.User, saltKey []byte) (string, error) {
	salt, err := encrypt(user.Salt, saltKey)
	if err != nil {
		return "", err
	}
	saltedPwd := saltPassword(user.Password, salt)

	hash, err := bcrypt.GenerateFromPassword(saltedPwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func saltPassword(password string, salt []byte) []byte {
	return append(salt, []byte(password)...)
}

// encrypt applies aes encryption to a string.
func encrypt(data string, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(data), nil), nil
}
