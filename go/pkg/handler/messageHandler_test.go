package handler

import (
	"go/pkg/authService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	id     = "1"
	config = middleware.JWTConfig{
		Claims:     &authService.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6OSwiZXhwIjoxNjM1NDg5MDA1fQ.U1xyuL1YrcOeHl-24_ehEW0eG4KxBoi41CXRCeri_BE"
)

func TestGetChat(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/getchat/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	middleware.JWTWithConfig(config)(GetChat)(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetChatsPreview(t *testing.T) {
	// Setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderAuthorization, middleware.DefaultJWTConfig.AuthScheme+" "+token)
	c := e.NewContext(req, rec)
	c.SetPath("/a/chatspreview")

	middleware.JWTWithConfig(config)(GetChatsPreview)(c)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
}
