package main

import (
	"errors"
	"github.com/joho/godotenv"
	"go-cloud-drive/handler"
	"go-cloud-drive/middleware"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var err error
	err = godotenv.Load()
	if err != nil {
		slog.Error("Fail to load .env file: " + err.Error())
		os.Exit(1)
	}

	port := os.Getenv("PORT")

	server := http.NewServeMux()

	server.HandleFunc("GET /hello", handler.Hello)

	slog.Info("Starting the server on port " + port + "...")

	err = http.ListenAndServe(":"+port, middleware.RequestLogger(server))

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("server closed")
	} else if err != nil {
		slog.Error("error starting server: %s\n", err.Error())
		os.Exit(1)
	}
}
