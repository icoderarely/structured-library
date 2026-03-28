package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/icoderarely/structured-library/internal/handler"
	"github.com/icoderarely/structured-library/internal/repository"
	"github.com/icoderarely/structured-library/internal/service"
)

func main() {
	// 1. Config
	port := getEnv("PORT", ":8080")

	// 2. Wire up dependencies
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	// 3. Routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", bookHandler.GetBooks)
	mux.HandleFunc("GET /books/{id}", bookHandler.GetByID)
	mux.HandleFunc("POST /books", bookHandler.CreateBook)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)

	// 4. Server with timeouts
	srv := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("starting server on %s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
