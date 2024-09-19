package main

import (
	"api_pgsql/configs"
	"api_pgsql/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server")
	r := chi.NewRouter()
	r.Get("/todos", handlers.List)
	r.Get("/todos/{id}", handlers.Get)
	r.Post("/todos", handlers.Create)
	r.Put("/todos/{id}", handlers.Update)
	r.Delete("/todos/{id}", handlers.Delete)
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
