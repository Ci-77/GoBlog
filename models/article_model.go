package models

import "gvb/models/ctype"

//定义文章结构体模型
type ArticleModel struct {
	MODEL
	Title        string         `gorm:"size:32" json:"title"`             //文章标题
	Abstract     string         ` json:"abstract"`                        //文章简介
	Content      string         ` json:"content"`                         //文章内容
	LookCount    int            ` json:"look_count"`                      //文章浏览量
	CommentCount int            ` json:"comment_count"`                   //文章评论量
	DiggCount    int            ` json:"digg_count"`                      //文章点赞量
	CollectCount int            ` json:"collect_count"`                   //文章收藏量
	Tag          []Tagmodel     `gorm:"many2many:article_tag_models" json:"tag"` //文章标签
	CommentList  []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`    //文章的评论列表
	User         User           `gorm:"foreignKey:UserID" json:"-"`       //文章作者
	UserID       uint           ` json:"user_id"`                         //用户id
	Category     string         `gorm:"size:20" json:"category"`          //文章分类
	Source       string         ` json:"source"`                          //文章来源
	Link         string         `json:"link"`                             //原文链接
	Banner       BannerModel    `gorm:"foreignKey:BannerID" json:"-"`     //文章封面
	BannerID     uint           ` json:"banner_id"`                       //封面id
	BannerPath   string         `json:"banner_path"`                      //跳转文章的封面
	Tages        ctype.Array    `gorm:"type:string;size" json:"tages"`    //跳转文章的标签
	NickName     string         `gorm:"size:42" json:"nick_name"`         //发布文章昵称
}
