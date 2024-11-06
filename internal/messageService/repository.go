package messageService

import (
	"errors"
	"gorm.io/gorm"
)

// MessageRepository - интерфейс для работы с базой данных
type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id uint, message Message) (Message, error)
	DeleteMessageByID(id uint) error
}

// messageRepository - структура, которая реализует интерфейс MessageRepository
type messageRepository struct {
	db *gorm.DB
}

// NewMessageRepository - конструктор для создания репозитория
func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

// CreateMessage - создаем сообщение в базе данных
func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

// GetAllMessages - получаем все сообщения из базы данных
func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

// UpdateMessageByID - обновляем сообщение по ID
func (r *messageRepository) UpdateMessageByID(id uint, message Message) (Message, error) {
	result := r.db.Model(&Message{}).Where("id = ?", id).Updates(message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

// DeleteMessageByID - удаляем сообщение по ID
func (r *messageRepository) DeleteMessageByID(id uint) error {
	// Удаляем сообщение с указанным ID
	result := r.db.Delete(&Message{}, id)

	// Если не было удалено ни одной строки, возвращаем ошибку
	if result.Error != nil {
		return result.Error
	}

	// Если запись не найдена и не удалена, возвращаем ошибку
	if result.RowsAffected == 0 {
		return errors.New("message not found")
	}

	// Если удаление прошло успешно, возвращаем nil
	return nil
}
