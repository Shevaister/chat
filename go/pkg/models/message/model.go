package message

import "time"

type Message struct {
	ID        int
	ChatID    int
	SenderID  int
	Text      string
	CreatedAt time.Time
}
