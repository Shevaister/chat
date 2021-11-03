package handler

import (
	"encoding/json"
	"go/pkg/authService"
	"go/pkg/db/repositoryInit/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	if len(request["password"].(string)) < 8 || len(request["login"].(string)) < 8 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	hashedPassword, err := authService.HashPassword(request["password"].(string))
	if err != nil {
		return err
	}
	success := user.Repository.SignUp(request["login"].(string), hashedPassword)
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

func SignIn(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}

	if len(request["password"].(string)) < 8 || len(request["login"].(string)) < 8 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	id, password, err := user.Repository.SignIn(request["login"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	success := authService.CheckPasswordHash(request["password"].(string), password)
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}

	_, _, avatar, err := user.Repository.FindUserById(id)
	if err != nil {
		return err
	}

	t := authService.GenerateToken(id)
	return c.JSON(http.StatusOK, map[string]string{"jwt": t, "avatar": avatar})
}
