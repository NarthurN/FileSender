package routes

import (
	"FileSender/FileSender/internal/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Обрабатываем маршрут для загрузки файлов
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/upload", handlers.UploadHandler)
	mux.HandleFunc("/uploadWithDrop", handlers.UploadWithDropHandler)
}
