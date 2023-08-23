package utils

import (
	"errors"
	"gvb/global"

	"github.com/dgrijalva/jwt-go/v4"
)

// 解析token
func ParseToken(tokenstr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.JWT.Secret)
	token, err := jwt.ParseWithClaims(tokenstr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {

		return MySecret, nil
	})
	if err != nil {
		global.Log.Error("token parse err...", err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
