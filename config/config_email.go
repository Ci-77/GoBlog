package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"`
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` //默认发件人的名字
	UserSSL           bool   `json:"use_ssl" yaml:"user_ssl"`                      //是否使用ssl
	UserTLS           bool   `json:"use_tls" yaml:"user_tls"`
}
