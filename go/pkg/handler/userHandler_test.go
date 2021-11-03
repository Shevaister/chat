package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	newLogin, _    = json.Marshal(map[string]interface{}{"login": "qwerty29"})
	newPassword, _ = json.Marshal(map[string]interface{}{"password": "qwerty21"})
	findUser, _    = json.Marshal(map[string]interface{}{"login": "qwerty21"})
)

func TestFindUser(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(findUser))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/find")

	middleware.JWTWithConfig(config)(FindUser)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestChangePassword(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newPassword))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/changepassword")

	middleware.JWTWithConfig(config)(ChangePassword)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestChangeLogin(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(newLogin))
	res := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, res)
	c.SetPath("/a/changelogin")

	middleware.JWTWithConfig(config)(ChangeLogin)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetID(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, res)
	c.SetPath("/a/id")

	middleware.JWTWithConfig(config)(GetID)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, res.Code)
}
