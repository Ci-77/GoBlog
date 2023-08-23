package router

import "gvb/api"

func (r RouterGroup) AdvertRouter() {
	advertapi := api.ApiGroupApp.AdvertApi
	r.POST("adverts", advertapi.AdvertCreateView)
	r.GET("adverts", advertapi.AdvertListView)
	r.PUT("adverts/:id", advertapi.AdvertUpdateView)
	r.DELETE("adverts", advertapi.AdvertRemoveView)
}
