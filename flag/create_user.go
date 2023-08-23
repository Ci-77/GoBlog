package flag

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/ctype"
	utils "gvb/utils/pwd"
)

// 写一个创建不用角色的函数
func CreateUser(permissions string) {
	var (
		userName   string
		nickName   string
		password   string
		repassword string
		email      string
	)
	//获得输入的数据
	fmt.Printf("请输入用户名:")
	fmt.Scanln(&userName)
	fmt.Printf("请输入用户昵称：")
	fmt.Scanln(&nickName)
	fmt.Printf("请输入密码：")
	fmt.Scanln(&password)
	fmt.Printf("请确认密码：")
	fmt.Scanln(&repassword)
	fmt.Printf("请输入邮箱：")
	fmt.Scanln(&email)
	// fmt.Println(userName,nickName,passward,repassward,email)
	//判断用户名是否存在，如果存在，直接登录
	//如果不存在，就继续创建用户步骤
	var user models.User
	//查找用户名，找到了就是存在的，没有错误
	err := global.DB.Take(&user, "user_name=?", userName).Error
	if err == nil {
		global.Log.Error("用户名已存在,请登录")
		return
	}
	//
	//检验两次密码
	if password != repassword {
		global.Log.Error("两次密码不一致,请重新输入密码")
		return
	}
	//将密码变为hash，然后入库
	hashword, _ := utils.HashPwd(password)
	//确定创建用户的角色
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	//确定创建角色的头像
	avatar := "/uploads/avater/default.jpg"
	//将创建的用户保存进数据库
	err = global.DB.Create(&models.User{
		UserName:   userName,
		NickName:   nickName,
		Password:   hashword,
		Email:      email,
		Role:       role,
		SignStatus: ctype.SigunEmail,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "清华",
	}).Error
	if err != nil {
		global.Log.Error(err)
	}
	global.Log.Infof("用户创建成功%s", userName)
}
