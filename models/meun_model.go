package models

import "gvb/models/ctype"

//菜单表
type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogan       string        `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                              //简介
	AbstractTime int           ` json:"abstract_time"`                                                                           //简介的切换时间
	Banners   []BannerModel `gorm:"many2many:menu_banner_models;joinForeignkey:MenuID;joinReference:ImageID" json:"banners"` //菜单的图片列表
	MenuTime     int           `json:"menu_time"`                                                                                //菜单图片法人切换时间
	Sort         int           `gorm:"size:10" json:"sort"`                                                                      //菜单的顺序
}
