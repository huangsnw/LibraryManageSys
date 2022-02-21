/*
 * @title: 这里写标题
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @BasePath: 这里写base path
 */
package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var MySecret = []byte("夏天夏天悄悄过去")

const TokenExpireDuration = time.Hour * 2

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

/**
 * @description:
 * @param {*}
 * @return {*}
 */
func GetToken(username string) string {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	res, _ := token.SignedString(MySecret)
	return res
}

/**
 * @description:
 * @param {string} tokenString
 * @return {*}
 */
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
