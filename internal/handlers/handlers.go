package handlers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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
	if r.Method == http.MethodPost {
		// Получаем все файлы
		r.ParseMultipartForm(10000)
		files := r.MultipartForm.File["files"]
		for _, fileHeader := range files {
			go func(fileHeader *multipart.FileHeader) {
				// Открываем файл
				file, err := fileHeader.Open()
				if err != nil {
					http.Error(w, "Ошибка при открытии файла", http.StatusBadRequest)
					return
				}
				defer file.Close()

				// Создаем файл на сервере
				dst, err := os.Create(fileHeader.Filename)
				if err != nil {
					http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
					return
				}
				defer dst.Close()

				// Копируем содержимое загружаемого файла
				if _, err := io.Copy(dst, file); err != nil {
					http.Error(w, "Ошибка при записи файла", http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "Файл %s успешно загружен.\n", fileHeader.Filename)
			}(fileHeader)
		}
	}
}

// uploadHandler обрабатывает загрузку файла
func UploadWithDropHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method == http.MethodPost {
		// Получаем все файлы
		r.ParseMultipartForm(10000)
		files := r.MultipartForm.File["filesDragAnDrope"]
		for _, fileHeader := range files {
			go func(fileHeader *multipart.FileHeader) {
				// Открываем файл
				file, err := fileHeader.Open()
				if err != nil {
					http.Error(w, "Ошибка при открытии файла", http.StatusBadRequest)
					return
				}
				defer file.Close()

				// Создаем файл на сервере
				dst, err := os.Create(fileHeader.Filename)
				if err != nil {
					http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
					return
				}
				defer dst.Close()

				// Копируем содержимое загружаемого файла
				if _, err := io.Copy(dst, file); err != nil {
					http.Error(w, "Ошибка при записи файла", http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "Файл %s успешно загружен.\n", fileHeader.Filename)
			}(fileHeader)
		}
	}
}
