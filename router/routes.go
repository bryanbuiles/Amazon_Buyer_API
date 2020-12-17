package router

import (
	"fmt"

	"github.com/bryanbuiles/tecnical_interview/api/consumer/gateway"
	"github.com/go-chi/chi"
)

// SetupData to setup the buyer, product and transaction data
func SetupData(app *chi.Mux) {
	fmt.Println("entro")
	star := gateway.Start()
	app.Route("/load", func(r chi.Router) {

		r.Get("/", star.ConsumerDateHandler)
	})
}
