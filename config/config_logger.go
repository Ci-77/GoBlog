package config
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`//log的前缀
	Director     string `yaml:"director"`
	LogInConsole bool   `yaml:"log-in-console"`      //是否显示行号
	ShowLine     bool   `yaml:"show_line"` //是否显示打印的路径
}