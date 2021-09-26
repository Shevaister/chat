package message

import (
	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func ProvideMessageRepostiory(DB *gorm.DB) MessageRepository {
	DB.AutoMigrate(&Message{})
	return MessageRepository{DB: DB}
}

func (m MessageRepository) GetAllChatMessages(id int) []Message {
	var messages []Message
	m.DB.Where("chat_id = ?", id).Find(&messages)
	return messages
}

func (m MessageRepository) GetLastChatMessage(id int) (Message, error) {
	var message Message
	err := m.DB.Where("chat_id = ?", id).Last(&message).Error
	return message, err
}

func (m *MessageRepository) SendMessage(chatId, userId int, text string) Message {
	message := Message{ChatID: chatId, SenderID: userId, Text: text}
	m.DB.Create(&message)
	return message
}
