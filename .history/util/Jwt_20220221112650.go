package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var MySecret = []byte("夏天夏天悄悄过去")

const TokenExpireDuration = time.Hour * 2

type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(username string) (string, error) {
	c := Token{
		Username: username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "my-project",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}
