package userser

import (
	redisser "gvb/service/redis_ser"
	jwt "gvb/utils/jwt"
	"time"
)

type UserService struct {
}
//设置过期时间，然后
func (UserService)Logout(claims *jwt.CustomClaims,token string) error {
exp:=claims.ExpiresAt
now:=time.Now()
diff:=exp.Time.Sub(now)//当前时间距离过期时间的时间差
return redisser.Logout(token,diff)

}