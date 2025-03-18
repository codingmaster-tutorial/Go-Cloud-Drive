package handler

import (
	"log/slog"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	slog.Info("Hello, world")
}
