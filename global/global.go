package global

import (
	"gvb/config"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//编写公共变量
var(
	Config *config.Config
	DB *gorm.DB
	Log *logrus.Logger
	Redis *redis.Client
)