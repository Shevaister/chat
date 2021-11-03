package wsService

import (
	"fmt"
	"go/pkg/db/repositoryInit/chat"
	"go/pkg/db/repositoryInit/message"
	"go/pkg/db/repositoryInit/social"
	"go/pkg/db/repositoryInit/user"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func init() {
	go handleMessages()
}

var Clients = make(map[int]*ClientConn)

type ClientConn struct {
	Conn     *websocket.Conn
	Requests int
}

// struct of the chat message that we get from client
type ChatMessage struct {
	SenderId   int    `json:"senderId"`
	ReceiverId int    `json:"receiverId"`
	Text       string `json:"text"`
}

var MessageChannel = make(chan ChatMessage, 1024)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func handleMessages() {
	for {
		msg := <-MessageChannel
		blocked := social.Repository.CheckBlocked(msg.ReceiverId, msg.SenderId)
		if blocked {
			Clients[msg.SenderId].Conn.WriteJSON(map[string]interface{}{"code": 0})
		} else {
			chatId := chat.Repository.GetChatID(msg.ReceiverId, msg.SenderId)
			sendedMessage := message.Repository.SendMessage(chatId, msg.SenderId, msg.Text)
			// send message to yourself (for correct time and id)
			if Clients[msg.SenderId] != nil {
				sendMessageWithChatPreview(Clients[msg.SenderId].Conn, sendedMessage.Text, chatId, sendedMessage.SenderID, msg.ReceiverId, sendedMessage.ID, sendedMessage.CreatedAt)
			}
			// send message to user
			if Clients[msg.ReceiverId] != nil {
				sendMessageWithChatPreview(Clients[msg.ReceiverId].Conn, sendedMessage.Text, chatId, sendedMessage.SenderID, sendedMessage.SenderID, sendedMessage.ID, sendedMessage.CreatedAt)
			}
		}
	}
}

func sendMessageWithChatPreview(client *websocket.Conn, text string, chatId, senderId, receiverId, messageId int, time time.Time) {
	_, senderLogin, senderAvatar, err := user.Repository.FindUserById(receiverId)
	if err != nil {
		fmt.Print(err)
	}
	// this data needed to make correct response in chat (if chat with this user open) and in chat preview
	err = client.WriteJSON(map[string]interface{}{"chatId": chatId, "messageId": messageId, "userAvatar": senderAvatar, "userLogin": senderLogin, "userId": senderId, "text": text, "time": time, "senderId": receiverId})
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		delete(Clients, senderId)
	}
}

func CloseConn(id int) {
	if Clients[id].Requests == 0 {
		Clients[id].Conn.Close()
		delete(Clients, id)
	} else {
		Clients[id].Requests--
	}
}
