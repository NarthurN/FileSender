package main

import (
	"fmt"
	"net/http"
	"time"

	"FileSender/FileSender/internal/routes"
)

func main() {
	router := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	routes.RegisterRoutes(router)

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на :8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
