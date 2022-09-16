package seeder

import (
	c "github.com/arioki1/alterra-agmc/day4/config"
	"github.com/arioki1/alterra-agmc/day4/models"
	"gorm.io/gorm"
	"log"
)

type seed struct {
	DB *gorm.DB
}

func NewUserSeeder() *seed {
	c.InitDB()
	return &seed{DB: c.DB}
}

func (s *seed) Seed() {
	users := []models.Users{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:     "test1",
			Email:    "test1@mail.com",
			Password: "test1",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:     "test2",
			Email:    "test2@mail.com",
			Password: "test2",
		},
	}
	if err := s.DB.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Println("success seed data users")
}

func (s *seed) Delete() {
	s.DB.Exec("DELETE FROM users")
}
