package models
//定义文章标签模型
type Tagmodel struct{
	MODEL
	Title string             `gorm:"size:16" json:"title"`    //标签名称
	Article []ArticleModel   `gorm:"many2many:article_tag_models" json:"-"` //关联该标签的文章列表
}