package router

import (
	"go/pkg/authService"
	"go/pkg/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	e.POST("/signin", handler.SignIn)
	e.POST("/signup", handler.SignUp)
	e.Static("/", "../public")
	e.GET("/ws/:Authorization", handler.Point)

	r := e.Group("/a")
	config := middleware.JWTConfig{
		Claims:     &authService.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.POST("/find", handler.FindUser)

	r.POST("/changepassword", handler.ChangePassword)
	r.POST("/changelogin", handler.ChangeLogin)
	r.POST("/changeavatar", handler.ChangeAvatar)

	r.POST("/add", handler.AddUser)
	r.POST("/block", handler.BlockUser)
	r.POST("/unfriend", handler.UnfriendUser)
	r.GET("/friends", handler.GetFriends)
	r.GET("/id", handler.GetID)

	r.GET("/chat/:id", handler.GetChat)
	r.GET("/chatspreview", handler.GetChatsPreview)

	return e
}
