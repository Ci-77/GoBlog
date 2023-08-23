package core

import (
	"fmt"
	"gvb/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	//gorm连接mysql初始化
	dsn := global.Config.Mysql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("数据库连接失败 err:%s\n", dsn))

	}
	return db
}
