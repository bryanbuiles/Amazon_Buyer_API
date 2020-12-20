package router

import (
	"github.com/bryanbuiles/tecnical_interview/api/v1/gateway"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

// Routes to setup the buyer, product and transaction data
func Routes() *chi.Mux {
	mux := chi.NewMux()
	star := gateway.Start()
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs
		cors.Handler(cors.Options{
			AllowedMethods: []string{"GET", "POST"},
			AllowedOrigins: []string{"*"},
		}),
	)
	mux.Get("/load", star.DataHandler)
	return mux
}
