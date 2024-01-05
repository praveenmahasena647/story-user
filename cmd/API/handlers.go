package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/praveenmahasena647/users/cmd/middleware"
	"github.com/praveenmahasena647/users/cmd/postgres"
	"github.com/praveenmahasena647/users/cmd/views"
)

func create(ctx echo.Context) error {
	var c = views.CreateAccount("CreateAccount")
	return c.Render(context.Background(), ctx.Response())
}

func createAndLogin(ctx echo.Context) error {
	var name, email, password = ctx.FormValue("email"), ctx.FormValue("name"), ctx.FormValue("password")
	var hashed, passwordErr = middleware.GeneratePassword([]byte(password))
	if passwordErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	var u = postgres.NewUser(name, email, string(hashed))
	fmt.Println(u)
	if c, err := u.Insert(); err != nil {
		fmt.Println(c)
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable)
	}
	return ctx.Redirect(http.StatusTemporaryRedirect, "/login")
}
