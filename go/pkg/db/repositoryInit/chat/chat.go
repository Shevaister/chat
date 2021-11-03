package chat

import (
	"go/pkg/db/connection"
	"go/pkg/repositories/chat"
)

var Repository chat.ChatRepository

func init() {
	Repository = chat.ProvideChatRepostiory(connection.DB)
}
