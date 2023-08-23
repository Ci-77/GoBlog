package articleapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

// 文章的修改
func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	id := c.Param("id")

	var cr ArticleListResponse
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailWithMsg("请求体绑定失败", c)
		return
	}
	var article models.ArticleModel
	err = global.DB.Take(&article, "id=?", id).Error
	if err != nil {
		response.FailWithMsg("此文章不存在", c)
		return
	}
	err=global.DB.Model(&article).Updates(map[string]any{
		"title":    cr.Title,
		"abstract": cr.Abstract,
		"content":  cr.Content,
		"category": cr.Category,
		"source":   cr.Source,
		"link":     cr.Link,
		"tages":    cr.Tages,
	}).Error
	if err!=nil {
		response.FailWithMsg("文章更新失败",c)
		return
	}
	response.OkWithMsg("文章编辑成功",c)
}
