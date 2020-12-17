package main

import (
	"github.com/bryanbuiles/tecnical_interview/router"
)

func main() {
	mux := router.SetupData()
	server := newServer(mux)
	server.Run()
}
