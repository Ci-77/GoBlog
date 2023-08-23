package advertapi

import (
	"gvb/models"
	"gvb/models/response"
	"gvb/service/common"
	"strings"

	"github.com/gin-gonic/gin"
)

//查询数据库所有的广告，然后保存在一个人列表里面
func (AdvertApi)AdvertListView(c *gin.Context)  {
	//先绑定结构体
	var cr models.PageInfo
	err:=c.ShouldBindQuery(&cr)
	if err!=nil {
		response.FailWithCode(response.ArguementError,c)
		return
	}
	//bool值的默认零值为false
	//根据请求头referer来判断是否包含admin
	referer:=c.GetHeader("Referer")
	ishow:=true
	if strings.Contains(referer,"admin") {
		//如果从admin来查的显示默认false
		ishow=false
	}
		//分页查询
	list,count,_:=common.ComList(models.AdvertModel{IsShow: ishow}, common.Option{
		PageInfo: cr,
		Debug: true,
	})
 response.OkwithList(list,count,c)
}