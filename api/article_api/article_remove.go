package articleapi

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

func(ArticleApi) ArticleRemoveView(c *gin.Context){
	//文章删除,定义一个删除列表，存储删除的id号，实现批量删除
	var cr models.RemoveRequest
	err:=c.ShouldBind(&cr)
	if err!=nil{
		response.FailWithMsg("删除参数绑定失败",c)
	}
	var articleRemoveList models.ArticleModel
	count:=global.DB.Find(&articleRemoveList,cr.IDList).RowsAffected
	if count==0 {
		response.FailWithMsg("此文章已删除",c)
	}
	err=global.DB.Delete(&articleRemoveList).Error
	if err!=nil {
		response.FailWithMsg("文章删除失败",c)
	}
	response.OkWithMsg(fmt.Sprintf("文章删除成功，共删除%d",count),c)
}