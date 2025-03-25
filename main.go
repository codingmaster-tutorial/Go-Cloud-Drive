package main

import (
	"errors"
	"github.com/joho/godotenv"
	"go-cloud-drive/handler"
	"go-cloud-drive/middleware"
	"go-cloud-drive/utils"
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

	db := utils.GetDB()
	if db == nil {
		slog.Error("Fail to connect database")
		os.Exit(1)
	}

	err = os.MkdirAll(os.Getenv("ROOT_DIR"), 0755)
	if err != nil {
		slog.Error("Fail to create root dir " + err.Error())
		os.Exit(1)
	}

	port := os.Getenv("PORT")

	server := http.NewServeMux()

	server.HandleFunc("GET /hello", handler.Hello)

	// TODO: Upload file
	server.HandleFunc("POST /file", handler.UploadFile)
	// TODO: Retrieve a list of files

	// TODO: Retrieve single file metadata

	// TODO: Edit file metadata

	// TODO: Delete file

	slog.Info("Starting the server on port " + port + "...")

	err = http.ListenAndServe(":"+port, middleware.RequestLogger(server))

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("server closed")
	} else if err != nil {
		slog.Error("error starting server: %s\n", err.Error())
		os.Exit(1)
	}
}
