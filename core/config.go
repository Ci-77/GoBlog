package core

import (
	"gvb/config"
	"gvb/global"
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)
const Configfile="settings.yaml"
//读取yaml文件的配置
func InitConf()  {
	c:=new(config.Config)
	yamlConf,err:=os.ReadFile(Configfile)
	if err!=nil {
		log.Fatalf("读取文件失败%s",err)
	}
	err=yaml.Unmarshal(yamlConf,c)
	if err!=nil {
log.Fatalf("config init err:%s\n",err)
	}
	log.Println("config init succes")
	global.Config=c
	
}
//写文件操作-就是将更新后的文件读取出来
func SetYaml() error {
	byteData,err:=yaml.Marshal(global.Config)
if err!=nil {
	return err
}
err=os.WriteFile(Configfile,byteData,fs.ModePerm)
if err!=nil {
	global.Log.Error(err)
	return err
}
global.Log.Info("配置文件修改成功")
return nil
}
