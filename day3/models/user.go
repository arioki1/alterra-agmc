package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	Token    string `json:"token" form:"token"`
}

func (u *Users) Validation() error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}
