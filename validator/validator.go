package validator

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type Validator struct{}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if len(username) > 30 {
		return errors.New("username must be no longer than 30 characters")
	}
	if match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username); !match {
		return errors.New("username can only contain letters, numbers, and underscores")
	}
	return nil
}

func (v *Validator) ValidatePassword(password []byte) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

func (v *Validator) HashPassword(password []byte) ([]byte, error) {
	// Use a proper password hashing algorithm here, such as bcrypt or scrypt
	// Example using bcrypt:
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
