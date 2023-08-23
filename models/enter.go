package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}
type PageInfo struct {
	Page  int    `form:"page"`  //当前页
	Key   string `form:"key"`
	Limit int    `form:"limit"`  //显示的分页数
	Sort  string `form:"sort"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}