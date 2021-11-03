package handler

import (
	"go/pkg/wsService"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Websocket(c echo.Context) error {
	// golang server send request to itself, because of we cant send anything but url params in ws connection and still need jwt authorization
	req, err := http.NewRequest("GET", "http://localhost"+os.Getenv("SERVER_PORT")+"/a"+"/id", nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	req.Header.Set("Authorization", "Bearer"+" "+c.Param("Authorization"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Logger().Error(err)
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

	// upgrading connection to websocket
	ws, err := wsService.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	//only 1 connection per user allowed
	if wsService.Clients[id] != nil {
		wsService.Clients[id].Requests++
		wsService.Clients[id].Conn.Close()
		wsService.Clients[id].Conn = ws
	} else {
		wsService.Clients[id] = &wsService.ClientConn{Conn: ws, Requests: 0}
	}
	defer wsService.CloseConn(id)

	for {
		var msg wsService.ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			c.Logger().Error(err)
			break
		}
		if msg.Text == "ping" || msg.ReceiverId == 0 {
			err := ws.WriteJSON(map[string]interface{}{"code": 1, "text": "pong"})
			if err != nil {
				c.Logger().Error(err)
				break
			}
			continue
		}
		msg.SenderId = id
		wsService.MessageChannel <- msg
	}
	return err
}
