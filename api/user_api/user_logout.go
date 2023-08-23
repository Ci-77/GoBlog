package userapi

import (
	"gvb/global"
	"gvb/models/response"
	"gvb/service"

	jwt "gvb/utils/jwt"

	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
//实现用户的注销操作
_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr UserRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
	}
	token:=c.Request.Header.Get("token")
	err=service.ServiceApp.UserService.Logout(claims,token)
	if err!=nil {
		global.Log.Error(err)
		response.FailWithMsg("注销失败",c)
	}
	response.FailWithMsg("注销成功",c)

}