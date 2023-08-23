package config

import "fmt"

//设置redis映射
type Redis struct {
	IP       string `json:"ip" yaml:"ip"`               //ip地址
	Port     int    `json:"port" yaml:"port"`           //端口
	Password string `json:"password" yaml:"password"`   //密码
	PoolSize int    `json:"pool_size" yaml:"pool_size"` //连接池
}
func (r Redis)Addr()string  {
	return fmt.Sprintf("%s:%d", r.IP,r.Port)//输出端口
}