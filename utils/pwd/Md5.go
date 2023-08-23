package utils

import (
	"crypto/md5"
	"encoding/hex"
)
//生成哈希数，通过验证哈希数来判断文件是否已经存在
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}