package repository

import (
	"gorm-crud-project/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetUserByIdAndEmail(id string, email string) (models.User, error) {
	var user models.User

	result := r.DB.Where("id=? and email =?", id, email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil

}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
