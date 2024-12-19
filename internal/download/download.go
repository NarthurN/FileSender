package download

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
)

func UploadFile(w http.ResponseWriter, r *http.Request, nameOfFilesForm string) {
	var wg sync.WaitGroup
	// Получаем все файлы
	r.ParseMultipartForm(10000)
	files := r.MultipartForm.File[nameOfFilesForm]
	for _, fileHeader := range files {
		wg.Add(1)
		go func(fileHeader *multipart.FileHeader) {
			defer wg.Done()
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
	wg.Wait()
}
