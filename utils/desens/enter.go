package desens

import "strings"

//设置邮箱和电话号码脱敏
//实现电话号码脱敏
func DesensitizationTel(tel string)string  {

	//15597293038
	//155****3038
	if len(tel)<11 {
		return ""
	}
	return tel[:3]+"****"+tel[7:]
}
//实现邮箱脱敏
func DesensitizationEmail(email string) string  {
	//3275571844@qq.com
	//3********@qq.com
	//将字符串按照指定的字符分割
	emlist:=strings.Split(email, "@")
	if  len(emlist)!=2{
		return ""
	}
	return emlist[0][:1]+"*****@"+emlist[1]
}