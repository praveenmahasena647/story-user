package api

import (
	"github.com/labstack/echo/v4"
	h "github.com/praveenmahasena647/users/cmd/helpers"
)

type APIserver struct {
	listenAddr string
}

func NewAPIserver(port string) *APIserver {
	return &APIserver{
		listenAddr: port,
	}
}

func (s *APIserver) Run() error {
	var e = echo.New()
	e.Use(h.Gurd)
	e.GET("/create", handleCreate)
	e.POST("/create", handleCreateAndLogin)
	e.GET("/login", handlelogin)
	e.POST("/login", handleLogins)
	e.GET("/home", handleHome)
	e.POST("/home", handleHome)
	e.POST("/newpost", handlepost)
	return e.Start(":8081") // should change later
}
