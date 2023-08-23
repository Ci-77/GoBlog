package menuapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

// 实现对菜单的更新操作
func (MenuApi) MenuUpdteView(c *gin.Context) {
	//实现菜单的更新操作
	//1.绑定参数
	var cr MenuRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
		return
	}
	//2.绑定id
	id := c.Param("id")
	var menumodel models.MenuModel
	err = global.DB.Take(&menumodel, id).Error
	if err != nil {
		response.FailWithMsg("菜单不存在", c)
		return
	}
	//对原来的内容进行清空
	global.DB.Model(&menumodel).Association("Banners").Clear()
	if len(cr.ImageList) > 0 {
		var bannerlist []models.MenuBannerModel
		for _, sort := range cr.ImageList {
			bannerlist = append(bannerlist, models.MenuBannerModel{
				MenuID:   menumodel.ID,
				BannerID: sort.ImageId,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerlist).Error
		if err != nil {
			response.FailWithMsg("创建图片失败...", c)
		}
	}
	//普通更新
	err = global.DB.Model(&menumodel).Updates(map[string]any{
		"menu_title":    cr.MenuTitle,
		"menu_title_en": cr.MenuTitleEn,
		"slogan":        cr.Slogan,
		"abstract":      cr.Abstract,
		"abstract_time": cr.AbstractTime,
		"menu_time":     cr.MenuTime,
		"sort":          cr.Sort,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("修改菜单失败", c)
		return
	}
	response.OkWithMsg("修改菜单成功", c)
}
