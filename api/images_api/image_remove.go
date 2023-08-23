package imagesapi

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

func (ImagesApi)ImageRemoveView(c *gin.Context) {
      var cr models.RemoveRequest
	  //绑定cr
	  err:=c.ShouldBindJSON(&cr)
	  if err!=nil {
		response.FailWithCode(response.ArguementError,c)
	  }
	  var imagelist []models.BannerModel
count:=global.DB.Find(&imagelist,cr.IDList).RowsAffected
if count==0 {
	response.FailWithMsg("文件不存在",c)
	return
}
//删除图片
global.DB.Delete(&imagelist)

response.OkWithMsg(fmt.Sprintf("共删除%d张图片", count),c)

}