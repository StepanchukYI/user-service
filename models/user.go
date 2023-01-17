package models

import (
	"time"

	"github.com/google/uuid"
)

type Identity uint

type UserStatus int8

const (
	STATUS_CREATED UserStatus = 0

	STATUS_APPROVED UserStatus = 10

	STATUS_DELETED UserStatus = 99
)

type User struct {
	Id         uuid.UUID

	Email string

	Status UserStatus

	CreatedAt time.Time
	UpdatedAt time.Time


}

type UserProfile struct {
	UserId uuid.UUID

	Name       string
	SecondName string

}