package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", handlerUserGet)
		r.Get("/users/{id}", handlerUserGetById)
		r.Post("/users", handlerUserPost)
		r.Delete("/users/{id}", handlerUserDelete)
		r.Put("/users/{id}", handlerUserPut)
	})

	return r
}

func handlerUserGet(w http.ResponseWriter, r *http.Request) {

}
func handlerUserGetById(w http.ResponseWriter, r *http.Request) {

}

func handlerUserPost(w http.ResponseWriter, r *http.Request) {

}

func handlerUserDelete(w http.ResponseWriter, r *http.Request) {

}

func handlerUserPut(w http.ResponseWriter, r *http.Request) {

}
