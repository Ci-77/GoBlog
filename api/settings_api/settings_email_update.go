package settingapi

import (
	"gvb/config"
	"gvb/core"
	"gvb/global"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsEmailUpdatView(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithCode(response.ArguementError, c)
		return
	}

	global.Config.Email = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg(err.Error(), c)
	}
	response.OkWith(c)
}