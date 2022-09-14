package database

import (
	"fmt"
	"github.com/arioki1/alterra-agmc/day2/config"
	"github.com/arioki1/alterra-agmc/day2/models"
)

func GetUsers() (*[]models.Users, error) {
	var users []models.Users

	if e := config.DB.
		Find(&users).Error; e != nil {
		return nil, e
	}
	return &users, nil
}

func CreateUsers(user *models.Users) (*models.Users, error) {
	if e := config.DB.
		Create(&user).
		Error; e != nil {
		return nil, e
	}
	return user, nil
}

func GetUserById(id *int) (*models.Users, error) {
	user := new(models.Users)
	if e := config.DB.
		Where("id = ?", id).
		First(user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUsers(user *models.Users) (*models.Users, error) {
	db := config.DB.
		Where("id = ?", user.ID).
		Updates(user).
		First(user)

	if db.RowsAffected < 1 {
		return nil, fmt.Errorf("row with id=%v  cannot be update because it doesn't exist", user.ID)
	} else if db.Error != nil {
		return nil, db.Error
	}

	return user, nil
}
