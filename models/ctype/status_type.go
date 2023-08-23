package ctype

import "encoding/json"
//定义登录状态
type SignStatus int

const (
	SignQQ     SignStatus = 1 //qq登录
	SignGithub SignStatus = 2 //github登录
	SigunEmail SignStatus = 3 //邮箱登录
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ登录"
	case SignGithub:
		str = "Github登录"
	case SigunEmail:
		str = "邮箱登录"
	default:
		str = "其他登录"
	}
	return str
}