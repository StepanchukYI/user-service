package repository

import (
	"database/sql"

	"github.com/StepanchukYI/user-service/internal/datastruct"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type UserRepositoryInterface interface {
	Create(user *datastruct.User) error
	GetByID(id uuid.UUID) (*datastruct.User, error)
	Update(id uuid.UUID, user *datastruct.User) error
	Delete(id uuid.UUID) error
}

type UserRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *datastruct.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(id uuid.UUID) (*datastruct.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = $1"
	var user datastruct.User
	err := r.db.QueryRow(query, id).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id uuid.UUID, user *datastruct.User) error {
	query := "UPDATE users SET email = $1, password = $2 WHERE id = $3"
	_, err := r.db.Exec(query, user.Email, user.Password, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
