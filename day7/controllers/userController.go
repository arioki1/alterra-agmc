package controllers

import (
	"fmt"
	"github.com/arioki1/alterra-agmc/day7/lib/database"
	"github.com/arioki1/alterra-agmc/day7/models"
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
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"status": "user not found",
				"user":   nil,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status": e.Error(),
				"user":   nil,
			})
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

	if id != getUserId(c) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "you can edit your own data only",
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

	result, e := database.UpdateUser(newUser)
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

func DeleteUserByIdControllers(c echo.Context) error {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Invalid params",
			"user":   nil,
		})
	}

	if id != getUserId(c) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "you can delete your own data only",
			"user":   nil,
		})
	}

	if e := database.DeleteUser(&id); e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": e.Error(),
			"user":   nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "deleted",
		"user":   nil,
	})
}

func LoginUsersControllers(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	users, e := database.LoginUsers(&user)

	if e != nil {
		if e.Error() == "record not found" {
			return echo.NewHTTPError(http.StatusBadRequest, "wrong email or password")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
		}
	}
	user.Password = ""
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})

}

func getUserId(c echo.Context) int {
	var id int
	userId := c.Get("userId")
	if userId != nil {
		i, err := strconv.Atoi(fmt.Sprintf("%v", userId))
		if err == nil {
			id = i
		}
	}
	return id
}
