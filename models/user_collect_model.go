package models

import "time"

//用户收藏表,记录用户什么时候收藏了什么文章
type UserCollect struct {
	MODEL
	UserID       uint         `gorm:"primarykey"`
	UserModel    User         `gorm:"foreignkey:UserID"`
	ArticleID    uint         `gorm:"primarykey"`
	ArticleModel ArticleModel `gorm:"foreignkey:ArticleID"`
	CreateAt     time.Time
}
