package ctype

import "encoding/json"
//定义角色模型
type Role int

const (
	PermissionAdmin       Role = 1 //管理员
	PermissionUser        Role = 2 //普通登录人
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //被禁用的用户

)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
//枚举
func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁用的用户"
	default:
		str = "其他用户"
	}
	return str
}