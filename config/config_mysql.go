package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` //日志等级
	Config  string  `yaml:"config"` //高级配置
}

func (m Mysql) Dsn() string { 
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port)+")/"+m.Db+"?"+"charset=utf8mb4&parseTime=True&loc=Local"
}