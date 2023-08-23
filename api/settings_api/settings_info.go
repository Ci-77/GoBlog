package settingapi

import (
	"gvb/global"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)
//查询接口
func (SettingsApi) SettingsInfoView(c *gin.Context) {
response.OkWithData(global.Config.SiteInfo,c)


}