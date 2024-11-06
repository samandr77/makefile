package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// DB - глобальная переменная для подключения к базе данных
var DB *gorm.DB

// InitDB - инициализация подключения к базе данных
func InitDB() {
	var err error
	// Строка подключения к базе данных
	dsn := "host=127.0.0.1 user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	fmt.Println("Успешно подключено к базе данных!")
}

// CloseDB - закрытие соединения с базой данных
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Ошибка закрытия базы данных: %v", err)
	}
	sqlDB.Close()
}
