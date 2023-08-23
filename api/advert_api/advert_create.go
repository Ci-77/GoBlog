package advertapi

import (
	"github.com/gin-gonic/gin"
	"gvb/global"
	"gvb/models"
	"gvb/models/response"
)

// AdvertCreateView 广告创建
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [get]

// 添加广告函数
type AdvertRequest struct {
	Title  string ` json:"title"  binding:"required" msg:"请输入标题"`   //显示的标题
	Href   string `json:"href" binding:"required" msg:"跳转链接非法"`     //跳转连接
	Images string `json:"images" binding:"required" msg:"图片地址非法"`   //图片
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否显示"` //是否显示
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
		return
	}
	//重复查询
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title=?", cr.Title).Error
	if err == nil {
		response.FailWithMsg("该广告已经存在", c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("添加广告失败", c)
		return
	}
	response.OkWithMsg("添加广告成功", c)
}
