package routes

import (
	"github.com/arioki1/alterra-agmc/day2/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	//Api User
	v1.GET("/users", controllers.GetUserControllers)
	v1.POST("/users", controllers.CreateUserControllers)

	return e
}
