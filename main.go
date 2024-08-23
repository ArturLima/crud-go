package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Arturlima/crud-go/api"
	repository "github.com/Arturlima/crud-go/db"
	"github.com/Arturlima/crud-go/models"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("All systems offline")
}

func run() error {

	db := make(map[string]models.User)

	repo := repository.NewUserRepository(db)

	handler := api.NewHandler(repo)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil

}
