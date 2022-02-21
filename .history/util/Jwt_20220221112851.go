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
