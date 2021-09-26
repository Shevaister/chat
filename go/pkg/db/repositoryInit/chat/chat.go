package chat

import (
	"go/pkg/db/connection"
	"go/pkg/models/chat"
)

var Repository chat.ChatRepository

func init() {
	Repository = chat.ProvideChatRepostiory(connection.DB)
}
