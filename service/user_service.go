package service

import (
	"encoding/json"
	"errors"
	"gorm-crud-project/kafka"
	"gorm-crud-project/models"
	"gorm-crud-project/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) GetUserByIdAndEmail(id string, email string) (models.User, error) {
	if id == "" || email == "" {
		return models.User{}, errors.New("id and email are required")
	}
	user, err := s.Repo.GetUserByIdAndEmail(id, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	if user.Name == "" {
		return nil, errors.New("name is required")
	}

	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	if user.Age < 18 {
		return nil, errors.New("age must be 18 or above")
	}

	createdUser, err := s.Repo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	message, err := json.Marshal(createdUser)
	if err != nil {
		return nil, err
	}

	err = kafka.PublishUserCreated(message)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
