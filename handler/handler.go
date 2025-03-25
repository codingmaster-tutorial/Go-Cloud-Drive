package handler

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	slog.Info("Hello, world")
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	f, fileHandler, err := r.FormFile("file")
	if err != nil {
		slog.Error("Fail to get file from form: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer f.Close()

	//slog.Info("file name is " + fileHandler.Filename)
	//slog.Info(fmt.Sprintf("file size is %d bytes", fileHandler.Size))

	newFilePath := filepath.Join(os.Getenv("ROOT_DIR"), fileHandler.Filename)
	newFile, err := os.Create(newFilePath)
	if err != nil {
		slog.Error("Fail to create new file " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, f)
	if err != nil {
		slog.Error("Fail to save file " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
