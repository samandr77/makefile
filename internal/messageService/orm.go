package messageService

import (
	"gorm.io/gorm"
)

// Message - модель сообщения
type Message struct {
	gorm.Model
	Content string `json:"content" gorm:"type:text;not null"`
}

// Структура MessageRepository должна будет использовать GORM для операций с базой данных
// Пример миграции структуры Message
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Message{})
}
