package articleapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/ctype"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

// 文章的创建
type ArticleRequest struct {
	Title        string         `json:"title"`  //标题
    Abstract     string         `json:"abstract"`//简介
    Content      string         `json:"content"`//内容
    Category     string         `json:"category"`//分类
    Source       string         `json:"source"`//文章来源
    Link         string         `json:"link"`//文章链接
    BannerPath   string         `json:"banner_path"`//封面图片路径
    Tages        []string       `json:"tages"`//标签
    UserID       uint           `json:"user_id"`//用户id
    BannerID     uint           `json:"banner_id"`//封面id
}

// 文章创建方法
func (ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailError(err, &cr, c)
		return
	}
//创建用户并且关联已有用户
	err=global.DB.Create(&models.ArticleModel{
		Title:        cr.Title,
        Abstract:     cr.Abstract,
        Content:      cr.Content,
        LookCount:    0,
        CommentCount: 0,
        DiggCount:    0,
        CollectCount: 0,
        Tag:          []models.Tagmodel{},
        CommentList:  []models.CommentModel{},
        User:         models.User{},
        UserID:       cr.UserID,
        Category:     cr.Category,
        Source:       cr.Source,
        Link:         cr.Link,
        Banner:       models.BannerModel{},
        BannerID:     cr.BannerID,
        BannerPath:   cr.BannerPath,
        Tages:        ctype.Array(cr.Tages),
	}).Error
	if err!=nil {
		response.FailWithMsg("文章创建失败",c)
		return
	}
	response.OkWithMsg("文章创建成功",c)
}
