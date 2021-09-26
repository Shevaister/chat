package message

import (
	"go/pkg/db/connection"
	"go/pkg/models/message"
)

var Repository message.MessageRepository

func init() {
	Repository = message.ProvideMessageRepostiory(connection.DB)
}
