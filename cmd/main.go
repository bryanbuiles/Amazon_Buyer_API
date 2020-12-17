package main

import (
	"github.com/bryanbuiles/tecnical_interview/router"
)

func main() {
	mux := router.Routes()
	server := newServer(mux)
	server.Run()
}
