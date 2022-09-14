package database

import (
	"github.com/arioki1/alterra-agmc/day2/config"
	"github.com/arioki1/alterra-agmc/day2/models"
)

func GetUsers() (*[]models.Users, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return &users, nil
}

func CreateUsers(user *models.Users) (*models.Users, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func GetUserById(id *int) (*models.Users, error) {
	user := new(models.Users)
	if e := config.DB.Where("id = ?", id).First(user).Error; e != nil {
		return nil, e
	}
	return user, nil
}
