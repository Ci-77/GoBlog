package userapi
//写一个qq登录

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/ctype"
	"gvb/models/response"
	qqlogin "gvb/plugins/qq_login"
	pwd  "gvb/utils/pwd"
	jwt  "gvb/utils/jwt"
	random "gvb/utils/random"
	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	//就是当扫完qq二维码以后，会出现一个URL，带着code
	//从前端拿到code
	code := c.Query("code")
	if code == "" {
		response.FailWithMsg("没有code", c)
		return
	}
	fmt.Println(code)
	qqinfo, err := qqlogin.NewQQLogin(code)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	openID := qqinfo.OpenID
	var user models.User
//查数据库有无以前登录过，没有就创建一个新的用户存入数据库
	err = global.DB.Take(&user, "token=?", openID).Error
	if err != nil {
		//没有查到token,就创建用户
		//将明文密码设置成哈希
		hashpwd,_:=pwd.HashPwd(random.RandomString(16))
		user=models.User{
			UserName: openID,
			NickName: qqinfo.Nickname,
			Password: hashpwd,//密码随机生成,并设置为密文
			Avatar:   qqinfo.Avater,
			Addr:     "内网",
			Token:    qqinfo.OpenID,
			IP:       c.ClientIP(),
			Role: ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err=global.DB.Create(&user).Error
		if err!=nil{
			global.Log.Fatal("创建失败")
			response.FailWithMsg("注册失败",c)
			return
		}

	}
	//登录操作
	token,err:=jwt.GenToken(jwt.JWTPayLoad{
		Nickname:user.NickName,
		UserID:user.ID,
		Role:int(user.SignStatus),
		})
		if err!=nil {
			global.Log.Fatalln("token生成失败",err)
			return
		}
		response.OkWithData(token,c)
		
}
