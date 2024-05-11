package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"uwe/handler"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Get("/customer/{id}", handler.Make(handler.HandleGetCustomer))
	router.Post("/upload", handler.Make(handler.HandleUpload))
	router.Post("/file", handler.Make(handler.HandleUpload))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("API server running", "addr", listenAddr)
	http.ListenAndServe(listenAddr, router)
}
