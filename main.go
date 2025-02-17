package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/CleysonPH/dreampic.ai/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Use(handler.WithUser)

	// server the static files
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))

	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/login", handler.Make(handler.HandleLoginIndex))

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application is running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEverything() error {
	return godotenv.Load()
}
