package imagesapi

import (
	"gvb/global"
	"gvb/models"
	"gvb/models/response"

	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   string `json:"id"`   //返回图片的id
	Path string `json:"path"` //返回图片的路径
	Name string `json:"name"` //返回图片的名称
}

// 实现图片列表
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imagelist []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imagelist)
	response.OkWithData(imagelist, c)
	}


