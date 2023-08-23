package articleapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

type ArticleListResponse struct {
	Title    string `json:"title"`    //文章标题
	Abstract string `json:"abstract"` //文章简介
	Content  string `json:"content"`  //文章内容
	Category string `json:"category"` //文章
	Source   string `json:"source"`   //文章分类
	Link     string `json:"link"`     //文章来源
	Tages    string `json:"tages"`    //文章标签
	Username string `json:"user_name"` //文章作者
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var articles []models.ArticleModel
	err := global.DB.Preload("User").Order("created_at DESC").Find(&articles).Error
	if err != nil {
		response.FailWithMsg("查询失败", c)
	}
	//前端请求体对应的响应结构体
	var articleListResponses []ArticleListResponse
	for _, article := range articles {
		articleListResponse := ArticleListResponse{
			Title:    article.Title,
			Abstract: article.Abstract,
			Content:  article.Content,
			Category: article.Category,
			Source:   article.Source,
			Username: article.User.UserName,
			Link: article.Link,
			
		}
		articleListResponses = append(articleListResponses, articleListResponse)
	}

	response.OkWithData(articleListResponses, c)
}
