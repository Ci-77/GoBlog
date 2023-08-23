package tagapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

//标签表的创建
type TagRequest struct {
	Title string `json:"title"`
}
func (TagApi)TagCreateView(c *gin.Context)  {
	var cr TagRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
		return
	}
	//重复查询
	var tag models.Tagmodel
	err = global.DB.Take(&tag, "title=?", cr.Title).Error
	if err == nil {
		response.FailWithMsg("该标签已经存在", c)
		return
	}
	err = global.DB.Create(&models.Tagmodel{
		Title:  cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("添加标签失败", c)
		return
	}
	response.OkWithMsg("添加标签成功", c)
}