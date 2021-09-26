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
	success := user.Repository.SignUp(request["login"].(string), request["password"].(string))
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
	id, err := user.Repository.SignIn(request["login"].(string), request["password"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	t := authService.GenerateToken(id)
	return c.JSON(http.StatusOK, t)
}
