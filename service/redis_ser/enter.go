package redisser

import (
	"gvb/global"
	utils "gvb/utils/pwd"
	"time"
)

// 对redis存储token的操作
// token的注销操作
const prefix = "logout_"
//注销逻辑：通过将token放到redis来实现注销操作
func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}
//根据是否在redis来判断是否注销成功
func CheckLogout(token string) bool {
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false
}
