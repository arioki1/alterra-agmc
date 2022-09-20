package middlewares

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func UserAuthMiddlewares() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			//get token from header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "you are not authorized")
			}

			//get bearer token
			token := strings.Split(authHeader, " ")[1]
			userId, e := ValidateToken(token)
			if userId == 0 || e != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "you are not authorized")
			}

			c.Set("userId", userId)
			return next(c)

		}
	}
}
