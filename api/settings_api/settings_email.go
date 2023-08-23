package settingapi

import (
	"gvb/global"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsEmailView(c *gin.Context) {
response.OkWithData(global.Config.Email, c)

}