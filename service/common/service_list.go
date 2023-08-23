package common

import (
	"gvb/global"
	"gvb/models"

)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	//查询图片列表
	
	//分页函数
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Sort=="" {
		option.Sort="created_at desc"//默认为按时间往前排desc 按时间往后排为asc
	}
	count = global.DB.Select("id").Find(&list).RowsAffected
	//分页操作
	//limit是分页操作；offset是   order对标order by -对结果集进行排序，默认为升序asc,可手动改为降序desc
	err = global.DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
