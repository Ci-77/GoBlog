package router

import (
	"gvb/api"
	"gvb/middleware"
)

func (r RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.POST("user", userApi.EmailLoginView)                                //进行email登录
	r.POST("login", userApi.QQLoginView)                                  //跳转到qq登录
	r.GET("user", middleware.JwtAuth(), userApi.UserListView)             //返回一个用户列表
	r.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView) //超级管理员修改用户权限
	r.PUT("user", middleware.JwtAuth(), userApi.UserUpdateView)           //个人修改用户的信息，昵称，头像，密码
	r.POST("logout", middleware.JwtAuth(), userApi.LogoutView)            //注销账户
	r.DELETE("user",middleware.JwtAdmin(),userApi.UserRemoveView)  //删除操作
	//管理员登录
}
