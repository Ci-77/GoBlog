package config
type UploadFile struct {
	Size int `json:"size" yaml:"size"`  //上传图片的大小
	Path string `json:"path" yaml:"path"` //图片上传的目录
}