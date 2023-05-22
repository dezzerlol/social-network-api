package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Generate(plainTextPass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainTextPass), 12)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Matches(plainTextPass string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(plainTextPass))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
