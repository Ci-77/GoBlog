package middleware

import (
	"gvb/models/ctype"
	"gvb/models/response"
	redis_ser "gvb/service/redis_ser"
	jwt "gvb/utils/jwt"

	"github.com/gin-gonic/gin"
)

//设置用户中间件
func JwtAuth() gin.HandlerFunc  {
	return func(c *gin.Context){
		token:=c.Request.Header.Get("token")
		if token=="" {
			response.FailWithMsg("未携带token",c)
			c.Abort()
			return
		}
	
		claims,err:=jwt.ParseToken(token)
		if err!=nil {
			response.FailWithMsg("token错误",c)
			c.Abort()
			return
		}
		if redis_ser.CheckLogout(token) {
			response.FailWithMsg("token已经失效",c)
			c.Abort()
			return
		}
		c.Set("claims",claims)
	}
		
	
}
//超级管理员
func JwtAdmin()gin.HandlerFunc  {
	return func(c *gin.Context){
		token:=c.Request.Header.Get("token")
		if token=="" {
			response.FailWithMsg("未携带token",c)
			c.Abort()
			return
		}
	
		claims,err:=jwt.ParseToken(token)
		if err!=nil {
			response.FailWithMsg("token错误",c)
			c.Abort()
			return
		}
		//判断是否在token中
	if redis_ser.CheckLogout(token) {
		response.FailWithMsg("token已经失效",c)
		c.Abort()
		return
	}
	//登录的用户
	if claims.Role!=int(ctype.PermissionAdmin){
		response.FailWithMsg("您当前没有权限",c)
		c.Abort()
		return
	}

		c.Set("claims",claims)
	}
		
}