package config

type SiteInfo struct {
	CreatedAt   string `yaml:"created_at" json:"createdat"`
	BeiAn       string `yaml:"bei_an" json:"bei_an"`
	Title       string `yaml:"title" json:"title"`
	QqImage     string `yaml:"qq_image" json:"qqimage"`
	Version     string `yaml:"version" json:"version"`
	Email       string `yaml:"email" json:"email"`
	WechatImage string `yaml:"wechat_image" json:"wechatimage"`
	Name        string `yaml:"name" json:"name"`
	Job         string `yaml:"job" json:"job"`
	Addr        string `yaml:"ddr" json:"addr"`
	Solgan      string `yaml:"solgan" json:"solgan"`
	SolganEn    string `yaml:"slogan_en" json:"solganen"`
	Web         string `yaml:"web" json:"web"`
}
