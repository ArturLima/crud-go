package api

import (
	"encoding/json"
	"net/http"

	repository "github.com/Arturlima/crud-go/db"
	"github.com/Arturlima/crud-go/models"
	"github.com/Arturlima/crud-go/utils"
	"github.com/go-playground/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(repo repository.IUserRepository) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", handlerUserGet)
		r.Get("/users/{id}", handlerUserGetById)
		r.Post("/users", handlerUserPost(repo))
		r.Delete("/users/{id}", handlerUserDelete)
		r.Put("/users/{id}", handlerUserPut)
	})

	return r
}

func handlerUserGet(w http.ResponseWriter, r *http.Request) {

}
func handlerUserGetById(w http.ResponseWriter, r *http.Request) {

}
func handlerUserPost(repo repository.IUserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			utils.SendJSON(
				w,
				utils.Response{Error: "Please provide FirstName LastName and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}
		if err := validator.New().Struct(user); err != nil {
			utils.SendJSON(
				w,
				utils.Response{Error: "Please provide FirstName LastName and bio for the user"},
				http.StatusBadRequest,
			)
			return
		}
		newUser, err := repo.Insert(&user)
		if err != nil {
			utils.SendJSON(
				w,
				utils.Response{Error: "There was an error while saving the user to the database"},
				http.StatusInternalServerError,
			)
			return
		}

		utils.SendJSON(w, utils.Response{Data: newUser}, http.StatusOK)
	}
}

func handlerUserDelete(w http.ResponseWriter, r *http.Request) {

}

func handlerUserPut(w http.ResponseWriter, r *http.Request) {

}
