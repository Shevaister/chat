package handler

import (
	"encoding/json"
	"go/pkg/authService"
	"go/pkg/db/repositoryInit/user"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	requestedUsers := user.Repository.FindUser(request["login"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	users := make([]map[string]interface{}, len(requestedUsers))

	for i := 0; i < len(requestedUsers); i++ {
		users[i] = map[string]interface{}{"id": requestedUsers[i].ID, "login": requestedUsers[i].Login, "avatar": requestedUsers[i].Avatar}
	}
	return c.JSON(http.StatusOK, users)
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

	if len(request["password"].(string)) < 8 {
		return c.JSON(http.StatusBadRequest, nil)
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

	if len(request["login"].(string)) < 8 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	status := user.Repository.ChangeLogin(id, request["login"].(string))
	if !status {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

// file name should be id + file extension to avoid collision, old photos should be deleted or rewrited
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
		return err
	}
	defer src.Close()

	_, _, userAvatar, err := user.Repository.FindUserById(id)
	if err != nil {
		return err
	}

	savePath := os.Getenv("IMAGE_SAVE_PATH")

	if userAvatar != "" {
		err := os.Remove(savePath + userAvatar)
		if err != nil {
			return err
		}
	}

	// Destination

	fileName := strconv.Itoa(id) + filepath.Ext(file.Filename)
	dst, err := os.Create(savePath + fileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	user.Repository.ChangeAvatar(id, fileName)

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
