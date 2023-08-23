package userapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/ctype"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

//修改用户数据
//修改用户的权限
type UserRole struct {
	Role ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`   //修改用户角色
	UserID uint `json:"user_id" binding:"required" msg:"用户id错误"`

}
func (UserApi)UserUpdateRoleView( c *gin.Context)  {
	//参数绑定
	var cr UserRole
	err:=c.ShouldBindJSON(&cr)
	if err!=nil{
		response.FailError(err,&cr, c)
		return
	}
	var user models.User
	err=global.DB.Take(&user,cr.UserID).Error
	if err!=nil{
		response.FailWithMsg("未找到用户id",c)
		return
	}
    err=global.DB.Model(&user).Update("role",cr.Role).Error
	if err!=nil {
		response.FailWithMsg("权限更新失败",c)
		return
	}
	response.OkWithMsg("权限更新成功",c)
}