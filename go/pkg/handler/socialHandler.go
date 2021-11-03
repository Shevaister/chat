package handler

import (
	"encoding/json"
	"go/pkg/authService"
	"go/pkg/db/repositoryInit/social"
	"go/pkg/db/repositoryInit/user"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func BlockUser(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	success := social.Repository.BlockUser(id, int(request["user2Id"].(float64)))
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func AddUser(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	success := social.Repository.AddUser(id, int(request["user2Id"].(float64)))
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func UnfriendUser(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	social.Repository.UnfriendUser(id, int(request["user2Id"].(float64)))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func GetFriends(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	var err error

	friends := social.Repository.GetFriends(id)

	res := make([]map[string]interface{}, len(friends))

	for i := 0; i < len(friends); i++ {
		res[i] = make(map[string]interface{})
		res[i]["id"], res[i]["login"], res[i]["avatar"], err = user.Repository.FindUserById(friends[i].User2ID)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, res)
}
