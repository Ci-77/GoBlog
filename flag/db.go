package flag

import (

	"gvb/global"
	"gvb/models"
)

//数据库表的迁移
func Makemigrations() {
	var err error
	global.DB.SetupJoinTable(&models.User{},"CollectsModels",&models.UserCollect{})
	global.DB.SetupJoinTable(&models.MenuModel{},"Bannners",&models.MenuBannerModel{})
	//生成四张表的表结构
	err=global.DB.Set("gorm:table_options","ENGINE=InnoDB").AutoMigrate(
		&models.MenuBannerModel{},
		&models.Tagmodel{},
		&models.MessageModel{},
		&models.UserCollect{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.BannerModel{},
		&models.User{},
		&models.LoginDataModel{},
		&models.AdvertModel{},
		&models.FadeBackModel{},



	)
	if err!=nil {
		global.Log.Error("[error] 生成数据库失败")
	return
	}
	global.Log.Info("[success] 生成数据库成功")

}