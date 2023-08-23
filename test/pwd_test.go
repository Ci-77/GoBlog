package test

import (
	"fmt"
	utils "gvb/utils/pwd"
	"testing"
)

//哈希密码的测试用例
func TestPwd(t *testing.T)  {
//生成哈希密码
	fmt.Println(utils.HashPwd("12234"))
}