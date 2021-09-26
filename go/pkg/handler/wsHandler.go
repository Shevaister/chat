package handler

import (
	"fmt"
	"go/pkg/db/repositoryInit/chat"
	"go/pkg/db/repositoryInit/message"
	"go/pkg/db/repositoryInit/user"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func init() {
	go handleMessages()
}

var clients = make(map[int]*websocket.Conn)

// chat message struct that we get from client
type ChatMessage struct {
	SenderId   int    `json:"senderId"`
	ReceiverId int    `json:"receiverId"`
	Text       string `json:"text"`
}

var messageChannel = make(chan ChatMessage, 1024)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func Point(c echo.Context) error {

	// golang server request to itself, because of we cant send anything but url params in ws connection and still need jwt authorization
	req, err := http.NewRequest("GET", "http://localhost:8000/a/id", nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer"+" "+c.Param("Authorization"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	resp.Body.Close()

	id, err := strconv.Atoi(string(respBody))

	if err != nil {
		c.Logger().Error(err)
		return err
	}

	c.Logger().Error(len(clients))
	if clients[id] != nil {
		clients[id].Close()
		//cet += 1
		//c.Logger().Error(cet)
		//delete(clients, id)
	}

	// upgrading connection to websocket
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		c.Logger().Error("error")
		//ws.Close()
		return err
	}
	defer ws.Close()

	clients[id] = ws

	defer delete(clients, id)

	for {
		var msg ChatMessage
		fmt.Print(0)
		err := ws.ReadJSON(&msg)
		if err != nil {
			c.Logger().Error(err)
			//ws.Close()
			break
		}
		msg.SenderId = id
		messageChannel <- msg
	}

	//return c.JSON(http.StatusBadGateway, nil)
	//delete(clients, id)
	return err
}

func handleMessages() {
	for {
		msg := <-messageChannel
		chatId := chat.Repository.GetChatID(msg.ReceiverId, msg.SenderId)
		sendedMessage := message.Repository.SendMessage(chatId, msg.SenderId, msg.Text)
		// send message to yourself (for correct time and id)
		if clients[msg.SenderId] != nil {
			sendMessageWithChatPreview(clients[msg.SenderId], sendedMessage.Text, chatId, sendedMessage.SenderID, sendedMessage.ID, sendedMessage.CreatedAt)
		}
		// send message user
		if clients[msg.ReceiverId] != nil {
			sendMessageWithChatPreview(clients[msg.ReceiverId], sendedMessage.Text, chatId, sendedMessage.SenderID, sendedMessage.ID, sendedMessage.CreatedAt)
		}
		//fmt.Println(msg)

	}
}

func sendMessageWithChatPreview(client *websocket.Conn, text string, chatId, senderId, messageId int, time time.Time) {
	_, senderLogin, senderAvatar, err := user.Repository.FindUserById(senderId)
	if err != nil {
		fmt.Println()
		fmt.Println(err)
		fmt.Println()
	}
	// this data needed to make correct response in chat (if chat with this user open) and in chat preview
	err = client.WriteJSON(map[string]interface{}{"chatId": chatId, "messageId": messageId, "senderAvatar": senderAvatar, "senderLogin": senderLogin, "senderId": senderId, "text": text, "time": time})
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, senderId)
	}
}
