package settingapi

import (
	"gvb/config"
	"gvb/core"
	"gvb/global"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

//修改接口
func (SettingsApi) SettingsUpdatInfoView(c *gin.Context) {
	var cr config.SiteInfo
	err:=c.ShouldBindJSON(&cr)
	if err!=nil {
		response.FailWithCode(response.ArguementError,c)
		return
	}

 global.Config.SiteInfo=cr
 err=core.SetYaml()
 if err!=nil {
	global.Log.Error(err)
	response.FailWithMsg(err.Error(),c)
	return
 }
response.OkWith(c)
}