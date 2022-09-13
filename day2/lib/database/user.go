package database

import (
	"github.com/arioki1/alterra-agmc/config"
	"github.com/arioki1/alterra-agmc/models"
)

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}
