package datastruct

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID

	Email    string
	Password []byte

	Status int8

	CreatedAt time.Time
	UpdatedAt time.Time
}
