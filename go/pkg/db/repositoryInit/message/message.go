package message

import (
	"go/pkg/db/connection"
	"go/pkg/repositories/message"
)

var Repository message.MessageRepository

func init() {
	Repository = message.ProvideMessageRepostiory(connection.DB)
}
