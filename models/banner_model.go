package models

import (
	"gvb/global"
	"gvb/models/ctype"
	"os"

	"gorm.io/gorm"
)

// 定义封面图片模型
type BannerModel struct {
	MODEL
	Path string          `json:"path"`                        //图片路径
	Hash string          `json:"hash"`                        //图片的hash,用于判断重复图片
	Name string          `gorm:"size:38" json:"name"`         //图片名称
	Type ctype.ImageType `gorm:"default:1" json:"image_type"` //图片保存的类型
}

// 在同一个事务中删除数据的钩子函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.Type == ctype.Local {
		err := os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
		}
		return err
	}
	return nil
}
//在同一个事务里的更新数据的钩子函数

