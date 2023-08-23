package utils

import (
	"gvb/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// 生成token
func GenToken(user JWTPayLoad) (string, error) {
	MySecret = []byte(global.Config.JWT.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.JWT.Expires))),
			Issuer:    global.Config.JWT.Issure,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
