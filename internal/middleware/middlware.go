package middleware

import (
	"net/http"
	"log"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логирование информации о запросе
		log.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

		// Добавление заголовка к ответу
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Вызов следующего обработчика
		next(w, r)
	}
}