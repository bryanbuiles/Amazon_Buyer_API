package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bryanbuiles/tecnical_interview/internal/logs"
	"github.com/go-chi/chi"
)

// MyServer ...
type MyServer struct {
	server *http.Server
}

// NewServer ...
func newServer(mux *chi.Mux) *MyServer {
	s := &http.Server{
		Addr:           ":3000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logs.Info("Api is running is port 3000")
	return &MyServer{s}
}

// Run ...
func (s *MyServer) Run() {
	log.Fatal(s.server.ListenAndServe())
}
