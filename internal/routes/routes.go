package routes

import (
	"FileSender/FileSender/internal/handlers"
	"FileSender/FileSender/internal/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Обрабатываем маршрут для загрузки файлов
	mux.HandleFunc("/", middleware.Logging(handlers.Index))
	mux.HandleFunc("/upload", middleware.Logging(handlers.UploadHandler))
	mux.HandleFunc("/uploadWithDrop", middleware.Logging(handlers.UploadWithDropHandler))
}
