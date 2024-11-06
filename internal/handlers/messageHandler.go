package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"structur/internal/messageService"
)

// Handler - структура обработчика с сервисом для работы с сообщениями
type Handler struct {
	Service *messageService.MessageService
}

// NewHandler - конструктор для создания нового обработчика
func NewHandler(service *messageService.MessageService) *Handler {
	return &Handler{Service: service}
}

// GetMessagesHandler - обработчик для получения всех сообщений
func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем все сообщения через сервис
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		// Если произошла ошибка, возвращаем код ошибки 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Устанавливаем заголовок контента как JSON
	w.Header().Set("Content-Type", "application/json")
	// Отправляем данные сообщений в ответ
	json.NewEncoder(w).Encode(messages)
}

// PostMessageHandler - обработчик для создания нового сообщения
func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messageService.Message
	// Декодируем JSON из тела запроса в структуру message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		// Если произошла ошибка при декодировании, возвращаем ошибку 400
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Создаем новое сообщение через сервис
	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		// Если произошла ошибка при создании, возвращаем ошибку 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Устанавливаем заголовок контента как JSON
	w.Header().Set("Content-Type", "application/json")
	// Отправляем созданное сообщение в ответ
	json.NewEncoder(w).Encode(createdMessage)
}
func (h *Handler) PatchMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	var message messageService.Message
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMessage, err := h.Service.UpdateMessageByID(uint(id), message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMessage)
}

// DeleteMessageHandler - обработчик для удаления сообщения по ID
func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteMessageByID(uint(id))
	if err != nil {
		if err.Error() == "message not found" {
			http.Error(w, "Message not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
