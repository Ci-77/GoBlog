package models

//消息模型定义
type MessageModel struct {
	MODEL
	SendUserID       uint   `gorm:"primarykey" json:"send_user_id"` //发送人id
	SendUserModel    User   `gorm:"foreignkey:SendUserID" json:"send_user_model"`
	SendUserNickName string `gorm:"size:42" json:"send_user_nick_name"`
	SendUserAvatar   string ` json:"send_user_avatar"`
	RevUserID        uint   `gorm:"primarykey" json:"re_user_id"` //接收人id

	RevUserModel    User   `gorm:"foreignkey:RevUserID" json:"-"`
	ReUserNickName string `gorm:"size:42" json:"re_user_nick_name"`
	ReUserAvatar   string ` json:"re_user_avatar"`
	IsRead         bool   `gorm:"default:false" json:"is_read"` //接收方是否查看
	Content        string ` json:"content"`                     //接收内容
}
