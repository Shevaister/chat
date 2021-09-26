package handler

import (
	"encoding/json"
	"fmt"
	"go/pkg/authService"
	"go/pkg/db/repositoryInit/user"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func FindUser(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	requestedUserId, requestedUserLogin, err := user.Repository.FindUser(request["login"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"id": requestedUserId, "login": requestedUserLogin})
}

func ChangePassword(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	user.Repository.ChangePassword(id, request["password"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func ChangeLogin(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	status := user.Repository.ChangeLogin(id, request["login"].(string))
	if !status {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func ChangeAvatar(c echo.Context) error {

	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println(err) //err
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("C:/jProject/chat/mysql/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	user.Repository.ChangeAvatar(id, file.Filename)

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func GetID(c echo.Context) error {
	cUser := c.Get("user").(*jwt.Token)
	claims := cUser.Claims.(*authService.JwtCustomClaims)
	id := claims.ID

	return c.String(http.StatusOK, strconv.Itoa(id))
}
