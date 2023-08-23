package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gvb/config"
	"gvb/global"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)
     //本函数主要解决上传七牛云的操作

//鉴权函数
func getToken( q config.QiNiu) string {
	accesskey:=q.AccessKey
	secrekey:=q.SecretKey
	bucket:=q.Bucket
	//storage包，提供了资源的上传，处理，资源处理等功能
	putPolicy:=storage.PutPolicy{
		Scope: bucket,
	}
	mac:=qbox.NewMac(accesskey,secrekey)
	upToken:=putPolicy.UploadToken(mac)
	
	
	return upToken
}
//获取上传配置
func GetCfj(q config.QiNiu) storage.Config  {
	cfg:=storage.Config{}
	zone,_:=storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone=&zone
	//是否使用httpas域名
	cfg.UseHTTPS=false
	//空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	//上传是否使用cdn加速
	cfg.UseCdnDomains=false
	return cfg
}
//实现文件的上传
func UploadImage(data []byte,imagename string,prefix string) (filepath string,err error)  {
	if !global.Config.QiNiu.Enable {
		return "",errors.New("请启用七牛云上传")
	}
	q:=global.Config.QiNiu
	if q.AccessKey==""||q.SecretKey=="" {
		return "",errors.New("请配置Accesskey和Secretkey")
	}
	if float64(len(data))/1024/1024>float64(q.Size) {
		return "",errors.New("文件大小超出上传限制")
	}
	upToken:=getToken(q)
	cfg:=GetCfj(q)
	formUploader:=storage.NewFormUploader(&cfg)
	ret:=storage.PutRet{}
	putExtra:=storage.PutExtra{
		Params:map[string]string{},
	}
	dataLen:=int64(len(data))
	//获取当前时间
	now:=time.Now().Format("20060102150405")
	key:=fmt.Sprintf("%s/%s__%s",prefix,now,imagename)
	err=formUploader.Put(context.Background(),&ret,upToken,key,bytes.NewReader(data),dataLen,&putExtra)
	if err!=nil {
   return "",err
	}
	return fmt.Sprintf("%s%s",q.CDN,ret.Key),nil
}