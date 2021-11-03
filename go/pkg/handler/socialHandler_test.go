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
	userId, _ = json.Marshal(map[string]interface{}{"user2Id": 4})
)

func TestBlockUser(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(userId))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/block")

	middleware.JWTWithConfig(config)(BlockUser)(c)

	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestAddUser(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(userId))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/add")

	middleware.JWTWithConfig(config)(AddUser)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUnfriendUser(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(userId))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/unfriend")

	middleware.JWTWithConfig(config)(UnfriendUser)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetFriends(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/friends")

	middleware.JWTWithConfig(config)(GetFriends)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}
