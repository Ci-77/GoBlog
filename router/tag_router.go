package router

import "gvb/api"

func (r RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	r.POST("tag", tagApi.TagCreateView)   //创建标签
	r.GET("tag", tagApi.TagListView)      //返回标签列表
	r.PUT("tag/:id", tagApi.TagUpdateView)    //更新标签列表
	r.DELETE("tag", tagApi.TagRemoveView) //删除标签列表

}
