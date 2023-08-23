package utils

import (


	"strings"
)

//判断是否存在于列表内
func InList(key string,list []string) bool {
for _, v := range list {
	if strings.EqualFold(key,v) {
		return true
	}
}
	return false
}
