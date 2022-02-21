package util

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(username string) string {
	token := Token{
		Username: username,
		jwt.StandardClaims{},
	}
}
