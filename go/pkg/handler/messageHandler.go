package handler

import (
	"go/pkg/authService"
	"go/pkg/db/repositoryInit/chat"
	"go/pkg/db/repositoryInit/message"
	"go/pkg/db/repositoryInit/user"

	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetChat(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	user2_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	_, login, avatar, err := user.Repository.FindUserById(user2_id)
	if err != nil {
		return err
	}

	chatId := chat.Repository.GetChatID(id, user2_id)

	messages := message.Repository.GetAllChatMessages(chatId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"user": map[string]interface{}{"id": user2_id, "login": login, "avatar": avatar}, "chatId": chatId, "chat": messages})
}

func GetChatsPreview(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	chats := chat.Repository.GetChatsID(id)

	chatPreviews := make([]map[string]interface{}, len(chats))

	for i := 0; i < len(chats); i++ {
		m, err := message.Repository.GetLastChatMessage(chats[i].ID)
		text := m.Text
		time := map[int]interface{}{0: m.CreatedAt, 1: ""}
		s := 0
		if err != nil {
			s = 1
		}
		messageId := m.ID

		senderId := 0

		if chats[i].User1ID != id {
			senderId = chats[i].User1ID
		} else {
			senderId = chats[i].User2ID
		}

		senderId, senderLogin, senderAvatar, err := user.Repository.FindUserById(senderId)

		if err != nil {
			return err
		}

		chatPreviews[i] = map[string]interface{}{"chatId": chats[i].ID, "messageId": messageId, "senderAvatar": senderAvatar, "senderLogin": senderLogin, "senderId": senderId, "text": text, "time": time[s]}
	}

	return c.JSON(http.StatusOK, chatPreviews)
}
