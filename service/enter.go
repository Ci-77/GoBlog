package service

import userser "gvb/service/user_ser"

type ServiceGroup struct {
	UserService userser.UserService
}
var ServiceApp=new(ServiceGroup)