package main

import (
	"log"

	"github.com/praveenmahasena647/users/cmd"
)

func main() {
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
