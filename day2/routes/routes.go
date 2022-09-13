package routes

import (
	"github.com/arioki1/alterra-agmc/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	v1.GET("/users", controllers.GetUserControllers)

	return e
}
