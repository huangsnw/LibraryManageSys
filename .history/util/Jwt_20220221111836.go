package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(username string) string {
	token := Token{
		Username: username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",
		},
	}
}
