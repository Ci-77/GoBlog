package userapi

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//删除函数操作
func (UserApi)UserRemoveView(c *gin.Context)  {
	var cr models.RemoveRequest
	err:=c.ShouldBind(&cr)
	if err!=nil {
		response.FailError(err,"关联失败",c)
	}
	//根据菜单的id来删除菜单
var userlist []models.User//声明一个存放menumodel的列表
count:=global.DB.Find(&userlist,cr.IDList).RowsAffected
if count==0 {
	response.FailWithMsg("用户不存在",c)
}
//事务
err=global.DB.Transaction(func (tx *gorm.DB)error  {
	//ToDo：删除用户,消息表，评论表，文章发布表，收藏表
	err=global.DB.Delete(&userlist).Error
if err!=nil {
	global.Log.Error(err)
	
	return err
}
return nil
})
if err!=nil {
	global.Log.Error(err)
	response.FailWithMsg("用户删除失败",c)
	return
}
response.OkWithMsg(fmt.Sprintf("用户删除成功,共删除%d个用户",count),c,)
}