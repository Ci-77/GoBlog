package imagesapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"
	"gvb/service/common"

	"github.com/gin-gonic/gin"
)

//实现图片列表
func (ImagesApi) ImageListView(c *gin.Context) {
//分页操作
 var cr models.PageInfo
 //参数绑定
 err := c.ShouldBindQuery(&cr)
 if err != nil {
 	response.FailWithCode(response.ArguementError, c)
		return }
 //分页查询
 list, count, err := common.ComList(models.BannerModel{}, common.Option{
 	PageInfo: cr,
 	Debug:    true,
	})
 if err!=nil {
		global.Log.Fatalf("分页查询失败")
 }
	 response.OkwithList(list,count, c )
	}
