package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password []byte) (hashed []byte, err error) {
	return bcrypt.GenerateFromPassword(password, 10)
}

func CompareHashAndPassword(hashed []byte, password []byte) (matched bool, err error) {
	err = bcrypt.CompareHashAndPassword(hashed, password)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		matched = false
		err = nil
		return
	}
	if err != nil {
		return
	}

	matched = true
	return
}
