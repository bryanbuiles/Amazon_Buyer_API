package main

import (
	"net/http"

	"github.com/bryanbuiles/tecnical_interview/internal/logs"
	"github.com/bryanbuiles/tecnical_interview/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {

	app := chi.NewRouter()

	app.Use(middleware.Recoverer)
	app.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedOrigins: []string{"*"},
	}))
	router.SetupData(app)

	logs.Info("Api is running is port 3000")
	err := http.ListenAndServe(":3000", app)

	if err != nil {
		logs.Error("Error runing the server at port 3000 " + err.Error())
	}
}
