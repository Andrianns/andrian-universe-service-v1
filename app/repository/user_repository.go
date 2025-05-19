package repository

import (
	"github.com/Andrianns/andrian-universe-service-v1/app/config"
	"github.com/Andrianns/andrian-universe-service-v1/app/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id uint) error
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) FindAll() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func (r *userRepo) FindByID(id uint) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return user, err
}

func (r *userRepo) Create(user models.User) (models.User, error) {
	err := config.DB.Create(&user).Error
	return user, err
}

func (r *userRepo) Update(user models.User) (models.User, error) {
	err := config.DB.Save(&user).Error
	return user, err
}

func (r *userRepo) Delete(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
