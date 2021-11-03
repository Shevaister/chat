package authService

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id int) (t string) {
	// Set custom claims
	claims := &JwtCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	return
}
