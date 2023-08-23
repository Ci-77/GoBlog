package flag

//根据命令行来创建用户

import (
	sys_flag "flag"

	"github.com/fatih/structs"
)

// 迁移表结构
type Option struct {
	DB   bool
	User string //-u admin /user
}

// 解析命令行参数
func Parse() Option {
	//根据命令行来迁移数据库-db
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	//解析命令行参数写入到注册的flag里面
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// 是否停止web项目
func IsWebStop(option Option) (f bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}
	}
	return f
}

// 根据命令行执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	//提示命令行
	sys_flag.Usage()

}
