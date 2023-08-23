package router

import "gvb/api"

//文章路由
func (r RouterGroup)ArticleRouter(){
articleApi:=api.ApiGroupApp.ArticleApi
r.POST("article",articleApi.ArticleCreateView)//文章创建
r.GET("article",articleApi.ArticleListView)//文章列表
r.PUT("article/:id",articleApi.ArticleUpdateView)//文章更新操作
r.DELETE("article",articleApi.ArticleRemoveView)//文章删除
}