package router

import (
	"gvb/api"

)


func (r RouterGroup)SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	//本地系统信息配置查看
	r.GET("", settingsApi.SettingsInfoView)
	r.PUT("",settingsApi.SettingsUpdatInfoView)
	
}