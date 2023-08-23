package test

import (
	"gvb/core"
	"gvb/global"
	"testing"
)

//redis连接的测试文件
func TestRedis(t *testing.T){
	core.InitConf()
	global.Log = core.InitLogger()
	global.Redis=core.InitRedis()

}