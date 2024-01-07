package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/praveenmahasena647/users/cmd/helpers"
	"github.com/praveenmahasena647/users/cmd/postgres"
	p "github.com/praveenmahasena647/users/cmd/postgres"
	"github.com/praveenmahasena647/users/cmd/views"
)

func handleCreate(ctx echo.Context) error {
	var c = views.CreateAccount("Create Account")
	return c.Render(context.Background(), ctx.Response())
}

func handleCreateAndLogin(ctx echo.Context) error {
	var email, username, password = ctx.FormValue("email"), ctx.FormValue("username"), ctx.FormValue("password")
	var hashed, hashErr = helpers.GeneratePassword(password)

	if hashErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "password hashing error")
	}
	var u = p.NewUser(username, email, hashed)
	if _, err := u.Insert(); err != nil {
		fmt.Println(err)
		return err
	}
	var JWT, JWTErr = helpers.GenerateJWT(u.EmailID)
	if JWTErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnt get JWT")
	}
	ctx.SetCookie(&http.Cookie{
		Name:  "user-jwt",
		Value: JWT,
	})
	ctx.Response().Header().Set("email", u.EmailID)
	return ctx.Redirect(http.StatusPermanentRedirect, "/home")
}

func handleHome(ctx echo.Context) error {
	var c = views.Home("home Page")
	return c.Render(context.Background(), ctx.Response())
}

func handlelogin(ctx echo.Context) error {
	var c = views.Login("loginPage")
	return c.Render(context.Background(), ctx.Response())
}

func handleLogins(ctx echo.Context) error {
	var email, password = ctx.FormValue("email"), ctx.FormValue("password")
	var u, userErr = postgres.FetchUser(email)
	if userErr != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user does not exist")
	}
	if err := helpers.PasswordErr(u.Password, password); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "wrong password")
	}
	var JWT, JWTErr = helpers.GenerateJWT(u.EmailID)
	if JWTErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "couldnt get JWT")
	}
	ctx.SetCookie(&http.Cookie{
		Name:  "user-jwt",
		Value: JWT,
	})
	return ctx.Redirect(http.StatusPermanentRedirect, "/home")
}

func handlepost(ctx echo.Context) error {
	var cookie, cookieErr = ctx.Cookie("user-jwt")
	if cookieErr != nil {
		return ctx.Redirect(http.StatusForbidden, "/login")
	}
	var emailID, JWTErr = helpers.DecodeJWT(cookie.Value)
	if JWTErr != nil {
		return ctx.Redirect(http.StatusForbidden, "/login")
	}
	var user = postgres.U()
	if err := user.FetchByEmail(emailID); err != nil {
		return ctx.Redirect(http.StatusNotFound, "/login")
	}
	var title, post = ctx.FormValue("title"), ctx.FormValue("story")
	var p = postgres.NewPost(title, post, user.ID)
	if err := p.Insert(); err != nil {
		return ctx.String(http.StatusInternalServerError, "couldn't Insert into posts")
	}
	return ctx.String(http.StatusOK, "New story has been added")
}

func handleShowAll(ctx echo.Context) error {
	var collection, err = postgres.FetchAll()
	fmt.Println(collection, err)
	return nil
}
