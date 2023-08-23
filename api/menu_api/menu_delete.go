package menuapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

//实现菜单的删除操作
func (MenuApi)MenuDeleteView(c *gin.Context)  {
	var cr models.RemoveRequest
	err:=c.ShouldBind(&cr)
	if err!=nil {
		response.FailError(err,"关联失败",c)
	}
	//根据菜单的id来删除菜单
var menulist []models.MenuModel//声明一个存放menumodel的列表
count:=global.DB.Find(&menulist,cr.IDList).RowsAffected
if count==0 {
	response.FailWithMsg("菜单不存在",c)
}
global.DB.Delete(&menulist)
response.OkWithMsg("菜单删除成功..",c)
}