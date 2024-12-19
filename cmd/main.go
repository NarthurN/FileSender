package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Обрабатываем маршрут для загрузки файлов
    http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/uploadWithDrop", uploadWithDropHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
    // Отправляем HTML-форму для загрузки файла
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    template, err := template.ParseFiles("templates/index.html")
    if err != nil {
        err = fmt.Errorf("Ошибка при парсинге шаблона %s", "index.html")
        log.Printf("Ошибка при парсинге шаблона %v", err)
        return
    }
    template.Execute(w, "")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получаем все файлы
        r.ParseMultipartForm(10000)
		files := r.MultipartForm.File["files"]
		for _, fileHeader := range files {
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
			fmt.Fprintf(w, "Файл %s успешно загружен.<br>", fileHeader.Filename)
		}
	}
}

// uploadHandler обрабатывает загрузку файла
func uploadWithDropHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method == http.MethodPost {
		// Получаем все файлы
        r.ParseMultipartForm(10000)
		files := r.MultipartForm.File["filesDragAnDrope"]
		for _, fileHeader := range files {
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
			fmt.Fprintf(w, "Файл %s успешно загружен.<br>", fileHeader.Filename)
		}
	} 
}
