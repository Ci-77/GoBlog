package messageapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

// 创建一对一聊天
type MessageRequest struct {
	SendUserID uint   `json:"send_user_id"`
	RevUserID  uint   `json:"re_user_id"`
	Content    string `json:"content"`
}

func (MessageApi) MessageCreateView(c *gin.Context) {
	//绑定请求
	var cr MessageRequest
	err := c.ShouldBind(&cr)
	if err != nil {
		global.Log.Fatalf("参数绑定失败")
	}
	var senUser, recvUser models.User
	err = global.DB.Take(&senUser, cr.SendUserID).Error
	if err != nil {
		response.FailWithMsg("发送人不存在", c)
		return
	}
	err = global.DB.Take(&recvUser, cr.RevUserID).Error
	if err != nil {
		response.FailWithMsg("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		RevUserID:        cr.RevUserID,
		Content:          cr.Content,
		SendUserNickName: senUser.NickName,
		SendUserAvatar:   senUser.Avatar,
		ReUserNickName:   recvUser.NickName,
		ReUserAvatar:     recvUser.Avatar,
	}).Error
	if err != nil {
		response.FailWithMsg("消息创建失败", c)
	}
	response.FailWithMsg("消息创建成功", c)
}
