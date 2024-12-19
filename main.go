package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    // Обрабатываем маршрут для загрузки файлов
    http.HandleFunc("/upload", uploadHandler)

    // Запускаем сервер на порту 8080
    fmt.Println("Сервер запущен на :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Ошибка:", err)
    }
}

// uploadHandler обрабатывает загрузку файла
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Проверка метода запроса
    if r.Method == http.MethodPost {
        // Получаем файл из формы
        file, _, err := r.FormFile("file")
        if err != nil {
            http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
            return
        }
        defer file.Close()

        // Создаем файл на сервере
        dst, err := os.Create("uploaded_file.pdf")
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

        // Успешное сообщение
        fmt.Fprintln(w, "Файл успешно загружен.")
    } else {
        // Отправляем HTML-форму для загрузки файла
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprint(w, `
            <form action="/upload" method="post" enctype="multipart/form-data">
                <input type="file" name="file" accept=".pdf">
                <input type="submit" value="Загрузить">
            </form>
        `)
    }
}
