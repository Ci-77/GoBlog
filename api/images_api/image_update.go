package imagesapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

//修改图片名称
type ImageUpdateResponce struct {
	ID   uint `json:"id" binding:"required" msg:"请输入文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名"`
}

func (ImagesApi) UpdateImageName(c *gin.Context) {
 var cr ImageUpdateResponce
 err:=c.ShouldBindJSON(&cr)
 if err!=nil {
	response.FailError(err,&cr,c)
	return
 }
 var imagemodel models.BannerModel
 //根据id来查询是否存在
 err=global.DB.Take(&imagemodel,cr.ID).Error
 if err!=nil {
	response.FailWithMsg("文件不存在",c)
	return

 }
 //更新数据库图片名称，并返回
err=global.DB.Model(&imagemodel).Update("name",cr.Name).Error
if err!=nil {
	response.FailWithMsg(err.Error(),c)
	return
}
response.OkWithMsg("图片名修改成功",c)
}
