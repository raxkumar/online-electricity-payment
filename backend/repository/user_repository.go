package repository

import (
	config "com.electricity.online/db"
	"com.electricity.online/models"
)

type UserRepository struct{}

func (ur *UserRepository) CreateUser(user *models.User) error {
	return config.DatabaseClient.Table("users").Create(&user).Error
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DatabaseClient.Table("users").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := config.DatabaseClient.Table("users").Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) UpdateUserByID(userID string, updatedUser *models.User) (*models.User, error) {
	existingUser, err := ur.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if err := config.DatabaseClient.Model(&existingUser).Updates(updatedUser).Error; err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (ur *UserRepository) GetUsersByZone(zoneName string) ([]models.User, error) {
	var users []models.User
	if err := config.DatabaseClient.Table("users").Where("zone_id = ?", zoneName).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
