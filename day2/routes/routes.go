package routes

import (
	"github.com/arioki1/alterra-agmc/day2/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	//Api Book
	v1.GET("/books", controllers.GetBooksControllers)
	v1.POST("/books", controllers.CreateBookControllers)
	v1.GET("/books/:id", controllers.GetBookByIdControllers)

	//Api User
	v1.GET("/users", controllers.GetUserControllers)
	v1.POST("/users", controllers.CreateUserControllers)
	v1.GET("/users/:id", controllers.GetUserByIdControllers)
	v1.PUT("/users/:id", controllers.UpdateUserByIdControllers)
	v1.DELETE("/users/:id", controllers.DeleteUserByIdControllers)

	return e
}
