package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)
//返回结构体中的msg
func GetValidMsg(err error, obj any) string {
	//使用的时候，需要传obj的指针
	getobj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	if errs,ok:=err.(validator.ValidationErrors);ok{
		//断言成功
		for _, e := range errs {
			//循环每一个err
			//根据报错字段名，获取结构体的具体参数
			if f,exits:=getobj.Elem().FieldByName(e.Field());exits{
            msg:=f.Tag.Get("msg")
			return msg
			}
		}
	}
	return err.Error()
}