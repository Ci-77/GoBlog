package userapi

//登录请求
import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"
	utils "gvb/utils/jwt"
	pwd "gvb/utils/pwd"

	"github.com/gin-gonic/gin"
)

type LoginRequst struct {
	Username string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// 邮箱登录方法实现
func (UserApi) EmailLoginView(c *gin.Context) {
	var cr LoginRequst
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
	}
	//对输入的用户名和密码进行验证
	var usermodel models.User
	err = global.DB.Take(&usermodel, "user_name=?", cr.Username).Error
	if err != nil {
		global.Log.Fatalf("用户不存在")
		response.FailWithMsg("用户名或密码错误", c)
		return
	}
	ischeck := pwd.CheckPwd(usermodel.Password, cr.Password)
	if !ischeck {
		global.Log.Fatalf("用户密码错误")
		response.FailWithMsg("用户密码错误", c)
		return
	}
	//登录成功，生成token
	token, err := utils.GenToken(utils.JWTPayLoad{
		Nickname: usermodel.NickName,
		UserID:   usermodel.ID,
		Role:     int(usermodel.Role), 
	})
	if err != nil {
		global.Log.Fatalln("token生成失败", err)
		return
	}
	response.OkWithData(token, c)
}
