package routes

import (
	"github.com/arioki1/alterra-agmc/day4/controllers"
	"github.com/arioki1/alterra-agmc/day4/middlewares"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")
	v1Auth := e.Group("/v1", middlewares.UserAuthMiddlewares())

	//Api Login
	v1.POST("/login", controllers.LoginUsersControllers)

	//Api Book
	v1.GET("/books", controllers.GetBooksControllers)
	v1.GET("/books/:id", controllers.GetBookByIdControllers)
	v1Auth.POST("/books", controllers.CreateBookControllers)
	v1Auth.PUT("/books/:id", controllers.UpdateBookByIdControllers)
	v1Auth.DELETE("/books/:id", controllers.DeleteBookByIdControllers)

	//Api User
	v1Auth.GET("/users", controllers.GetUserControllers)
	v1Auth.GET("/users/:id", controllers.GetUserByIdControllers)
	v1.POST("/users", controllers.CreateUserControllers)
	v1Auth.PUT("/users/:id", controllers.UpdateUserByIdControllers)
	v1Auth.DELETE("/users/:id", controllers.DeleteUserByIdControllers)

	return e
}
