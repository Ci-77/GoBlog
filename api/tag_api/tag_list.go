package tagapi

import (
	"gvb/models"
	"gvb/models/response"
	"gvb/service/common"

	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	//先绑定结构体
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArguementError, c)
		return
	}
	//分页查询
	list, count, _ := common.ComList(models.Tagmodel{}, common.Option{
		PageInfo: cr,

	})
	response.OkwithList(list, count, c)
}