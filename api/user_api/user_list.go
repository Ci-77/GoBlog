package userapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/ctype"
	"gvb/models/response"
	"gvb/service/common"
	jwt "gvb/utils/jwt"
	desens "gvb/utils/desens"

	"github.com/gin-gonic/gin"
)

//用户列表页
func (UserApi)UserListView(c *gin.Context)  {
	//捕获token来看是否是管理员账号
_claims,_:=c.Get("claims")
claims:=_claims.(*jwt.CustomClaims)
	//实现分页查询
	var page models.PageInfo
	err:=c.ShouldBindQuery(&page)
	if err!=nil {
		response.FailWithCode(response.ArguementError,c)
		return
	}
	var users []models.User
	list, count, err := common.ComList(models.User{}, common.Option{
		PageInfo: page,
	   })
	if err!=nil {
		   global.Log.Fatalf("分页查询失败")
	}
	for _, user := range list {
		if claims.Role!=int(ctype.PermissionAdmin) {
			//不是管理员，就显示游客
			user.UserName=""
		}
		user.Tel=desens.DesensitizationTel(user.Tel)
		user.Email=desens.DesensitizationEmail(user.Email)

		users=append(users,user)
	}
	
		response.OkwithList(users,count, c )

}