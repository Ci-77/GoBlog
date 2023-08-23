package main

import (
	"fmt"
	"gvb/core"
	 "gvb/flag"
	"gvb/global"
	"gvb/router"
	_"gvb/docs"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth
func main() {
	//读取文件的方法
	core.InitConf()
	fmt.Println(global.Config)
	//调用封装的log
	global.Log = core.InitLogger()
	//gorm连接mysql的方法
	global.DB = core.InitGorm()
    global.Redis=core.InitRedis()
	//数据库表的迁移
	 option := flag.Parse()
 if flag.IsWebStop(option) {
	 	flag.SwitchOption(option)
	 	return
	 }
	addr := global.Config.System.Add()
	// 调用gin
	router := router.InitGin()
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
