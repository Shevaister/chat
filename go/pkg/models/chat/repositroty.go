package chat

import (
	"gorm.io/gorm"
)

type ChatRepository struct {
	DB *gorm.DB
}

func ProvideChatRepostiory(DB *gorm.DB) ChatRepository {
	DB.AutoMigrate(&Chat{})
	return ChatRepository{DB: DB}
}

func (c ChatRepository) GetChatID(id1, id2 int) int {
	var chat Chat
	if c.DB.Where("user1_id = ? AND user2_id = ? or user1_id = ? AND user2_id = ?", id1, id2, id2, id1).Take(&chat).Error != nil {
		chat.User1ID = id1
		chat.User2ID = id2
		c.DB.Create(&chat)
		return chat.ID
	}
	return chat.ID
}

func (c ChatRepository) GetChatsID(id int) []Chat {
	var chats []Chat
	c.DB.Where("user1_id = ? OR user2_id = ?", id, id).Find(&chats)
	return chats
}
