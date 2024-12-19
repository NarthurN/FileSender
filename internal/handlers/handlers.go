package handlers

import (
	"FileSender/FileSender/internal/download"
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Отправляем HTML-форму для загрузки файла
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка 500", http.StatusInternalServerError)
		log.Printf("Ошибка при парсинге шаблона templates/index.html")
		return
	}
	template.Execute(w, "")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		download.UploadFile(w, r, "files")
	}
}

// uploadHandler обрабатывает загрузку файла
func UploadWithDropHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	switch r.Method {
	case http.MethodPost:
		download.UploadFile(w, r, "filesDragAnDrope")
	}
}
