package chat

import (
	"go/pkg/models/chat"

	"gorm.io/gorm"
)

type ChatRepository struct {
	DB *gorm.DB
}

func ProvideChatRepostiory(DB *gorm.DB) ChatRepository {
	DB.AutoMigrate(&chat.Chat{})
	return ChatRepository{DB: DB}
}

// finds current chat id (when user opens chat with someone)
func (c ChatRepository) GetChatID(id1, id2 int) int {
	var chat chat.Chat
	if c.DB.Where("user1_id = ? AND user2_id = ? OR user1_id = ? AND user2_id = ?", id1, id2, id2, id1).Take(&chat).Error != nil {
		chat.User1ID = id1
		chat.User2ID = id2
		c.DB.Create(&chat)
		return chat.ID
	}
	return chat.ID
}

// finds all chats id where user participate (for chat previews)
func (c ChatRepository) GetChatsID(id int) []chat.Chat {
	var chats []chat.Chat
	c.DB.Where("user1_id = ? OR user2_id = ?", id, id).Find(&chats)
	return chats
}
