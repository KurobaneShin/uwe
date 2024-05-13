package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"uwe/db"
	"uwe/handler"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db := db.Create()
	uploadHandler := handler.NewUploadHandler(db)
	customerHandler := handler.NewCustomerHandler(db)

	router := chi.NewMux()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	router.Post("/customer", handler.Make(customerHandler.HandleCreateCustomer))

	router.Get("/customer/{id}", handler.Make(handler.HandleGetCustomer))
	router.Post("/file", handler.Make(uploadHandler.HandleCreateFileUpload))
	router.Post("/file/{id}", handler.Make(uploadHandler.HandleFileUpload))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("API server running", "addr", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
