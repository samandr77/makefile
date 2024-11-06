package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"structur/internal/database"
	"structur/internal/handlers"
	"structur/internal/messageService"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	defer database.CloseDB() // Закрыть соединение с базой данных после завершения работы приложения

	// Миграция моделей
	err := messageService.Migrate(database.DB)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	// Создание репозитория и сервиса
	repo := messageService.NewMessageRepository(database.DB)
	service := messageService.NewMessageService(repo)

	// Создание хендлеров
	handler := handlers.NewHandler(service)

	// Настройка маршрутов
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/messages", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/messages/{id}", handler.PatchMessageHandler).Methods("PATCH")
	router.HandleFunc("/messages/{id}", handler.DeleteMessageHandler).Methods("DELETE")

	// Запуск HTTP-сервера
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", router)
}
