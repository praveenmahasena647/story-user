package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Gurd(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		return next(c)
	}
}
