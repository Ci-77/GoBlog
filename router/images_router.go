package router

import "gvb/api"

//上传图片路由
func (r RouterGroup) ImagesRouter() {
	imageApi := api.ApiGroupApp.ImagesApi
	//上传图片以及查看
	//创建
	r.POST("images", imageApi.ImageUploadView)
	//查询
	r.GET("images", imageApi.ImageListView)
	r.GET("images_names", imageApi.ImageNameListView)
	//删除
	r.DELETE("images", imageApi.ImageRemoveView)
	//修改
	r.PUT("images", imageApi.UpdateImageName)
}
