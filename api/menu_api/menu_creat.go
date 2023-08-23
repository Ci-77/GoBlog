package menuapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/ctype"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

type ImageSortList struct {
	ImageId uint `json:"image_id"`
	Sort    int  `json:"sort"`
}
type MenuRequest struct {
	MenuTitle    string          `json:"menu_title"`    //菜单的中文名称
	MenuTitleEn  string          `json:"menu_title_en"` //菜单的英文名称
	Slogan       string          `json:"slogan"`        //
	Abstract     ctype.Array     `json:"abstract"`
	AbstractTime int             `json:"abstract_time"` //切换的时间，单位都为S
	MenuTime   int             `json:"menu_time"`   //切换的时间，单位都为S
	Sort         int             `json:"sort"`          //菜单的序号
	ImageList    []ImageSortList `json:"image_list"`    //实现图片添加的排序
}

// 实现菜单的创建
func (MenuApi) MenuCreateView(c *gin.Context) {
	//绑定模型
	var cr MenuRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
		return
	}

	menumodel := models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		MenuTime: cr.MenuTime,
		Sort:       cr.Sort,
	}
	err = global.DB.Create(&menumodel).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("菜单添加失败", c)
		return
	}
	if len(cr.ImageList) == 0 {
		response.OkWithMsg("菜单添加成功", c)
		return
	}
	//写第三张入库的操作
	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageList {
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menumodel.ID,
			BannerID: sort.ImageId,
			Sort:     sort.Sort,
		})
	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("菜单图片关联失败", c)
		return
	}

}
