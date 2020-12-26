package main

import (
	"github.com/bryanbuiles/tecnical_interview/internal/database"
	"github.com/bryanbuiles/tecnical_interview/router"
)

func main() {
	database.SetpUpSheme()
	mux := router.Routes()
	server := newServer(mux)
	server.Run()
}
