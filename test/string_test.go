package test

import (
	"fmt"
	random "gvb/utils/random"
	"testing"
)

func TestString(t *testing.T) {
	//随机生成16位的字符串
	s:=random.RandomString(16)
	fmt.Println(s)
}