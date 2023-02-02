package repository

import (
	"github.com/StepanchukYI/user-service/internal/datastruct"
	"github.com/StepanchukYI/user-service/storage"
	"github.com/google/uuid"
)

type InternalUserRepository struct {
	storage *storage.SafeMap[uuid.UUID, *datastruct.User]
}

func NewInternalUserRepository() *InternalUserRepository {
	return &InternalUserRepository{
		storage.New[uuid.UUID, *datastruct.User](),
	}
}

func (in *InternalUserRepository) Create(user *datastruct.User) error {
	in.storage.Insert(user.Id, user)

	return nil
}

func (in *InternalUserRepository) GetByID(id uuid.UUID) (*datastruct.User, error) {
	return in.storage.Get(id)
}

func (in *InternalUserRepository) Has(id uuid.UUID) bool {
	return in.storage.Has(id)
}

func (in *InternalUserRepository) Update(id uuid.UUID, user *datastruct.User) error {
	return in.storage.Update(id, user)
}

func (in *InternalUserRepository) Delete(id uuid.UUID) error {
	return in.storage.Delete(id)
}
