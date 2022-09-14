package controllers

import (
	"github.com/arioki1/alterra-agmc/day2/models"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var lastBookId int
var books []models.Book

func GetBooksControllers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func CreateBookControllers(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid JSON body",
			"users":  nil,
		})
	}
	lastBookId++
	book.Id = lastBookId
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	if err := book.Validation(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": err.Error(),
			"book":   nil,
		})
	}

	books = append(books, book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "created",
		"books":  book,
	})
}
