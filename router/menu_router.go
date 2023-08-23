package router

import "gvb/api"

func (r RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	r.POST("menu",menuApi.MenuCreateView)
	r.GET("menu",menuApi.MenuListView)
	r.PUT("menu/:id",menuApi.MenuUpdteView)
	r.DELETE("menu",menuApi.MenuDeleteView)
}