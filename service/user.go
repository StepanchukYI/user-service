package service

import (
	"github.com/StepanchukYI/user-service/models"
	"github.com/StepanchukYI/user-service/repository"
	"github.com/StepanchukYI/user-service/validator"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	userRepo  repository.UserRepository
	validator validator.Validator
}

func (s *userService) CreateUser(user *models.User) error {

	// Perform validation
	if err := s.validator.ValidateUsername(user.Email); err != nil {
		return err
	}
	if err := s.validator.ValidatePassword(user.Password); err != nil {
		return err
	}

	// Hash the password
	hashedPassword, err := s.validator.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return nil
	//return s.userRepo.Create(user)
}

func (s *userService) GetUserByID(id uuid.UUID) (*models.User, error) {
	//user, err := s.userRepo.GetByID(id)
	//if err != nil {
	//	return nil, err
	//}

	// Perform any necessary processing or manipulation on the returned user before returning it
	return nil, nil
	//return user, nil
}

func (s *userService) UpdateUser(user *models.User) error {
	// Perform validation
	if err := s.validator.ValidateUsername(user.Email); err != nil {
		return err
	}
	if err := s.validator.ValidatePassword(user.Password); err != nil {
		return err
	}

	// Hash the password
	hashedPassword, err := s.validator.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return nil
	//return s.userRepo.Update(user)
}

func (s *userService) DeleteUser(id uuid.UUID) error {
	return s.userRepo.Delete(id)
}
