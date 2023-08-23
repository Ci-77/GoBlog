package userapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"
	jwt "gvb/utils/jwt"
	pwd "gvb/utils/pwd"

	"github.com/gin-gonic/gin"
)

// 用户修改密码
// 用户修改自己的昵称
// 用户修改自己的头像
type UserRequest struct {
	Nickname string `json:"nickname"`
	Newpwd   string `json:"newpwd"`
	Avatar   string `json:"avatar_id"`
	Oldpwd   string `json:"oldpwd"`
}

func (UserApi) UserUpdateView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr UserRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
	}
	//传入用户修改的内容
	var user models.User
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		response.FailWithMsg("用户不存在", c)
	}
	//将写入的密码和旧密码进行对比
	if !pwd.CheckPwd(user.Password, cr.Oldpwd) {
		response.FailWithMsg("密码错误", c)
		return
	}
	//将新密码生成hash密码
	hashNewPwd, _ := pwd.HashPwd(cr.Newpwd)
	err = global.DB.Model(&user).Updates(map[string]any{
		"nick_name":  cr.Nickname,
		"password":  hashNewPwd,
		"avatar": cr.Avatar,
	}).Error
	if err != nil {
		response.FailWithMsg("用户信息更新失败", c)
		return
	}
	response.OkWithMsg("用户信息更新成功", c)
	//
}
