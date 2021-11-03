package message

import (
	mm "go/pkg/models/message" //mm = message model

	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

func ProvideMessageRepostiory(DB *gorm.DB) MessageRepository {
	DB.AutoMigrate(&mm.Message{})
	return MessageRepository{DB: DB}
}

func (m MessageRepository) GetAllChatMessages(id int) []mm.Message {
	var messages []mm.Message
	m.DB.Where("chat_id = ?", id).Find(&messages)
	return messages
}

func (m MessageRepository) GetLastChatMessage(id int) (mm.Message, error) {
	var message mm.Message
	err := m.DB.Where("chat_id = ?", id).Last(&message).Error
	return message, err
}

func (m *MessageRepository) SendMessage(chatId, userId int, text string) mm.Message {
	message := mm.Message{ChatID: chatId, SenderID: userId, Text: text}
	m.DB.Create(&message)
	return message
}
