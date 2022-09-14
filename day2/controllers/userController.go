package controllers

import (
	"github.com/arioki1/alterra-agmc/day2/lib/database"
	"github.com/arioki1/alterra-agmc/day2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserControllers(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserControllers(c echo.Context) error {
	var user models.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid JSON body",
			"users":  nil,
		})
	}

	if err := user.Validation(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": err.Error(),
			"user":   nil,
		})
	}

	users, e := database.CreateUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "created",
		"users":  users,
	})
}

func GetUserByIdControllers(c echo.Context) error {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid params",
			"user":   nil,
		})
	}

	users, e := database.GetUserById(&id)
	if e != nil {
		if e.Error() == "record not found" {
			return echo.NewHTTPError(http.StatusBadRequest, "user not found")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   users,
	})
}

func UpdateUserByIdControllers(c echo.Context) error {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid params",
			"user":   nil,
		})
	}

	var user models.Users
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid JSON body",
			"users":  nil,
		})
	}
	newUser := new(models.Users)
	newUser.ID = uint(id)
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.Password = user.Password

	result, e := database.UpdateUsers(newUser)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": e.Error(),
			"user":   nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   result,
	})
}
