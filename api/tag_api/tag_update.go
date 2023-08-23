package tagapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

//更新标签表
func (TagApi)TagUpdateView(c *gin.Context)  {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
		return
	}
	//重复查询
	var tag models.Tagmodel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		response.FailWithMsg("该标签不存在", c)
		return
	}
	
	err = global.DB.Model(&tag).Updates(map[string]any{
		"title":  cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("更新标签失败", c)
		return
	}
	response.OkWithMsg("更新标签成功", c)
}