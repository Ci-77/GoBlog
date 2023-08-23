package tagapi

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	//根据数据库内的id进行删除
	//1.绑定参数
	var cr models.RemoveRequest
	//绑定cr
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArguementError, c)
	}
	var taglist []models.Tagmodel
	count := global.DB.Find(&taglist, cr.IDList).RowsAffected
	if count == 0 {
		response.FailWithMsg("标签不存在", c)
		return
	}
	//删除标签
	global.DB.Delete(&taglist)

	response.OkWithMsg(fmt.Sprintf("共删除%d个标签", count), c)

}
