package repository

import (
	"testing"

	"github.com/StepanchukYI/user-service/internal/datastruct"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternalUserRepository(t *testing.T) {
	repo := NewInternalUserRepository()

	userId, err := uuid.NewUUID()
	assert.Nil(t, err)

	user := &datastruct.User{
		Id:     userId,
		Email:  "some@email.com",
		Status: 0,
	}

	err = repo.Create(user)
	assert.Nil(t, err)

	userFromMemory, err := repo.GetByID(userId)
	assert.Nil(t, err)

	assert.Equal(t, user, userFromMemory)
}
