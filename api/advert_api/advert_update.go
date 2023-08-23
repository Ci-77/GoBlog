package advertapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

//更新广告函数
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
		return
	}
	//重复查询
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		response.FailWithMsg("该广告不存在", c)
		return
	}
	//由于传的是类型是bool，而引用了后台的
	//结构体转map第三方包的使用
	err = global.DB.Model(&advert).Updates(map[string]any{
		"title":  cr.Title,
		"href":   cr.Href,
		"images": cr.Images,
		"is_Show": cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("编辑广告失败", c)
		return
	}
	response.OkWithMsg("编辑广告成功", c)
}
