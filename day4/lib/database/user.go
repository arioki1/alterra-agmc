package database

import (
	"fmt"
	"github.com/arioki1/alterra-agmc/day4/config"
	"github.com/arioki1/alterra-agmc/day4/middlewares"
	"github.com/arioki1/alterra-agmc/day4/models"
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

func UpdateUser(user *models.Users) (*models.Users, error) {
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

func DeleteUser(id *int) error {
	db := config.DB.
		Where("id = ?", id).
		Delete(&models.Users{})

	if db.RowsAffected < 1 {
		return fmt.Errorf("row with id=%v  cannot be delete because it doesn't exist", *id)
	} else if db.Error != nil {
		return db.Error
	}

	return nil
}

func LoginUsers(user *models.Users) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ? ", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
