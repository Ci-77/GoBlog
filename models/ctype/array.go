package ctype

import (
	"database/sql/driver"
	"strings"
)

//将标签存在数组内
type Array []string
func (t *Array) Scan(value interface{}) error {
	v,_:=value.([]byte)
	if string(v)=="" {
		*t=[]string{}

		return nil
	}
	*t=strings.Split(string(v), "\n")
	return nil
}
//将数字转化为值
func(t Array) Value()(driver.Value,error)  {
	return strings.Join(t,"\n"),nil
}
