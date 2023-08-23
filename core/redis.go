package core

import (
	"context"
	"gvb/global"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

//连接redis
func InitRedis() *redis.Client  {
return ConnectRedisDB(0)
}
func ConnectRedisDB(db int) *redis.Client  {
	redisconf:=global.Config.Redis
	rdb:=redis.NewClient(&redis.Options{
		Addr: redisconf.Addr(),
		Password: redisconf.Password,//密码
		DB: 0,
		PoolSize: redisconf.PoolSize,//连接池大小
	})
	_,cancel:=context.WithTimeout(context.Background(),500*time.Millisecond)
	defer cancel()
	_,err:=rdb.Ping().Result()
	if err!=nil {
		logrus.Errorf("redis连接失败%s",err)
		return  nil
	}
	return rdb

}