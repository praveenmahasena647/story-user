package cmd

import (
	api "github.com/praveenmahasena647/users/cmd/API"
	"github.com/praveenmahasena647/users/cmd/postgres"
)

func Start() error {
	if err := postgres.Connect(); err != nil {
		return err
	}
	var s = api.NewServer(":42069")
	return s.Run()
}
