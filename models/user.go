package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserStatus int8

const (
	STATUS_CREATED UserStatus = 0

	STATUS_APPROVED UserStatus = 10

	STATUS_DELETED UserStatus = 99
)

type User struct {
	Id uuid.UUID

	Email    string
	Password []byte

	Status UserStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) CreatePassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = hash
	return nil
}

type UserProfiles struct {
	Id uuid.UUID

	UserId uuid.UUID

	Name       string
	SecondName string
}
