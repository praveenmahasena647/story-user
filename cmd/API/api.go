package api

import (
	"github.com/labstack/echo/v4"
	m "github.com/praveenmahasena647/users/cmd/middleware"
)

type APIserver struct {
	listenAddr string
}

func NewServer(port string) *APIserver {
	return &APIserver{
		listenAddr: port,
	}
}

func (s *APIserver) Run() error {
	var e = echo.New()

	e.Use(m.Gurd)
	e.GET("/create", create)
	e.POST("/create", createAndLogin)

	return e.Start(s.listenAddr)
}
