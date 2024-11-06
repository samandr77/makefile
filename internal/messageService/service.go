package messageService

// MessageService - структура для работы с сообщениями
type MessageService struct {
	repo MessageRepository
}

// NewMessageService - конструктор для создания сервиса
func NewMessageService(repo MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// CreateMessage - вызываем метод репозитория для создания сообщения
func (s *MessageService) CreateMessage(message Message) (Message, error) {
	return s.repo.CreateMessage(message)
}

// GetAllMessages - вызываем метод репозитория для получения всех сообщений
func (s *MessageService) GetAllMessages() ([]Message, error) {
	return s.repo.GetAllMessages()
}

// UpdateMessageByID - вызываем метод репозитория для обновления сообщения
func (s *MessageService) UpdateMessageByID(id uint, message Message) (Message, error) {
	return s.repo.UpdateMessageByID(id, message)
}

// DeleteMessageByID - вызываем метод репозитория для удаления сообщения
func (s *MessageService) DeleteMessageByID(id uint) error {
	return s.repo.DeleteMessageByID(id)
}
