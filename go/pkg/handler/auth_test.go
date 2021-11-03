package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON, _ = json.Marshal(map[string]interface{}{"password": "qwerty21", "login": "qwerty23"})
)

func TestSignUp(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signup")

	//Assertions
	if assert.NoError(t, SignUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignIn(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/signin")

	//Assertions
	if assert.NoError(t, SignIn(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, "sft", rec.Body.String())
	}
}
