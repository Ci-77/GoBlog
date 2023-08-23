package config

type JWT struct {
	Secret  string `json:"secret" yaml:"secret"`  //密钥
	Expires int    `json:"expires" yaml:"expires"` //过期时间
	Issure  string `json:"issure" yaml:"issure"` //颁发人
}
