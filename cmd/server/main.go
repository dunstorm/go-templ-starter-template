package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dunstorm/go-templ-starter-template/config"
	"github.com/dunstorm/go-templ-starter-template/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	port := config.Port

	r.Get("/", handlers.NewHomeHandler().ServeHTTP)

	ipAndPort := fmt.Sprintf("127.0.0.1:%s", port)
	log.Println("Starting server on", ipAndPort)
	if err := http.ListenAndServe(ipAndPort, r); err != nil {
		log.Fatal(err)
	}
}
