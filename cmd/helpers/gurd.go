package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var notAllowed map[string]bool = map[string]bool{
	"/home":    true,
	"/newpost": true,
}

var withcookieNotAllowed map[string]bool = map[string]bool{
	"/login":  true,
	"/create": true,
}

func Gurd(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var path = c.Request().RequestURI
		var _, cookieErr = c.Cookie("user-jwt")
		if cookieErr != nil && notAllowed[path] {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		if cookieErr == nil && withcookieNotAllowed[path] {
			return c.Redirect(http.StatusTemporaryRedirect, "/home")
		}
		return next(c)
	}
}
