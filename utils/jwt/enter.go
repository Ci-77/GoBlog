package utils

import "github.com/dgrijalva/jwt-go/v4"

//jwt中的payload数据
type JWTPayLoad struct {
	Username string `json:"username"` //用户名称
	Nickname string `json:"nickname"` //用户昵称
	Role     int    `json:"role"`     //用户的角色 1.管理员 2.普通用户 3.游客
	UserID   uint   `json:"user_id"`  //用户id
}
var MySecret []byte  //秘钥
type CustomClaims struct {
	JWTPayLoad//用户信息结构体
	jwt.StandardClaims  //
}
