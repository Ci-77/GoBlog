package test

import (
	"fmt"
	"gvb/core"
	"gvb/global"
	jwt "gvb/utils/jwt"
	"testing"
)

// jwt的测试文件
func TestJwt(t *testing.T) {

	core.InitConf()
	global.Log = core.InitLogger()
	token,err:=jwt.GenToken(jwt.JWTPayLoad{
		UserID:   1,
		Username: "阿蔡",
		Role:     1,
		Nickname: "xxxx",
	})
	fmt.Println(token, err)
  claims,err:=jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IumYv-iUoSIsIm5pY2tuYW1lIjoieHh4eCIsInJvbGUiOjEsInVzZXJfaWQiOjEsImV4cCI6MTY5MjA4NTY4MS44OTU2ODMsImlzcyI6IjEyMzQifQ.lQv-og13PYagjW0yF2hlMqIonrf2eP9-64kYTmMYrjA")
 fmt.Println(claims,err)
}
