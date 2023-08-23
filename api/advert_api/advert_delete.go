package advertapi

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)
// AdvertRemoveView 删除广告
//	@Summary		删除广告
//	@Description	删除广告
//	@Tags			广告管理
//	@Accept			json
//	@Produce		json
//	@Param			id	body		AdvertRequest	true	"表示多个参数"
//	@Success		200	{object}	response.Resonse{}
//	@Failure		400	{object}	response.Resonse{}
//	@Failure		404	{object}	response.Resonse{}
//	@Failure		500	{object}	response.Resonse{}
//	@Router			/api/adverts [delete]
//广告的删除
func(AdvertApi) AdvertRemoveView(c *gin.Context)  {
	//根据数据库内的id进行删除
//1.绑定参数
var cr models.RemoveRequest
//绑定cr
err:=c.ShouldBindJSON(&cr)
if err!=nil {
  response.FailWithCode(response.ArguementError,c)
}
var advertlist []models.AdvertModel
count:=global.DB.Find(&advertlist,cr.IDList).RowsAffected
if count==0 {
response.FailWithMsg("广告不存在",c)
return
}
//删除图片
global.DB.Delete(&advertlist)

response.OkWithMsg(fmt.Sprintf("共删除%d个广告", count),c)

}

