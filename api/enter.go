package api

import (
	advertapi "gvb/api/advert_api"
	articleapi "gvb/api/article_api"
	imagesapi "gvb/api/images_api"
	menuapi "gvb/api/menu_api"
	messageapi "gvb/api/message_api"
	settingapi "gvb/api/settings_api"
	tagapi "gvb/api/tag_api"
	userapi "gvb/api/user_api"
)

type ApiGroup struct {
	SettingsApi settingapi.SettingsApi
	ImagesApi  imagesapi.ImagesApi
	AdvertApi advertapi.AdvertApi
	MenuApi  menuapi.MenuApi
	UserApi   userapi.UserApi
	ArticleApi articleapi.ArticleApi
	TagApi tagapi.TagApi
	MessageApi messageapi.MessageApi
}
//分配新的内存空间
var ApiGroupApp=new(ApiGroup)