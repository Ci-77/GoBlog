package models

//菜单图片模型
type MenuBannerModel struct {
	MenuID      uint       `json:"menu_id"`     //
	MenuModel   MenuModel  `gorm:"foreignkey:MenuID"`
	BannerID    uint       `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignkey:BannerID"`
	Sort        int        `gorm:"size:10" json:"sort"`
}

