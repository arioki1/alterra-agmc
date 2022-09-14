package models

import (
	"fmt"
	"time"
)

type Book struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Isbn      string    `json:"isbn"`
	Writer    string    `json:"writer"`
}

func (u *Book) Validation() error {
	if u.Title == "" {
		return fmt.Errorf("title is required")
	}
	if u.Isbn == "" {
		return fmt.Errorf("isbn is required")
	}
	if u.Writer == "" {
		return fmt.Errorf("writer is required")
	}
	return nil
}
