package router

import (
	"gvb/api"

	
)

func (r RouterGroup) MessageRouter() {
messageApi:=api.ApiGroupApp.MessageApi
r.POST("message",messageApi.MessageCreateView)//创建消息
}