package imagesapi

import (
	"fmt"
	"gvb/global"
	"gvb/models"
	"gvb/models/ctype"
	"gvb/models/response"
	"gvb/plugins/qiniu"
	utils "gvb/utils/pwd"

	"io"
	"io/fs"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

var (
	//图片上传白名单
	WhiteList = []string{
		".jpg",
		".png",
		".ico",
		".svg",
		".gif",
		".webp",
		".tiff",
		".jpeg",
	}
)

type UploadImagesResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //携带的消息
}

// 上传多个图片
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//上传单个文件 c.FormFile()
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}
	filelist, ok := form.File["images"]
	if !ok {
		response.FailWithMsg("不存在文件", c)
		return
	}
	//判断路径是否存在
	basepath := global.Config.UploadFile.Path
	_, err = os.ReadDir(basepath)
	if err != nil {
		//如果没有读到文件路径，创建文件
		err = os.MkdirAll(basepath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	//路径不存在创建路径
	var reslist []UploadImagesResponse
	//遍历保存图片的数组
	for _, file := range filelist {
		filename := file.Filename
		//读文件名的后缀
		suffix := path.Ext(filename)
		//调用函数，判断是否在定义的列表内
		if !utils.InList(suffix, WhiteList) {
			reslist = append(reslist, UploadImagesResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue
		}
		//写一个函数判断是否存在于列表内
		filepath := path.Join(basepath, file.Filename)
		//判断上传图片的大小
		size := float64(file.Size) / float64(1024*1024)
		if size > float64(global.Config.UploadFile.Size) {
			reslist = append(reslist, UploadImagesResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过限制,图片大小为：%.2fMB,设定大小为%dMB", size, global.Config.UploadFile.Size),
			})
			continue
		}
		fileobj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileobj)
		if err!=nil {
			global.Log.Fatalf("读取文件的hash失败")
		}
		imageHash := utils.Md5(byteData)
		//去数据库查图片是否存在
		var bannermodel models.BannerModel
		err=global.DB.Take(&bannermodel,"hash=?",imageHash).Error
		if err==nil {
			//找到了
			reslist = append(reslist, UploadImagesResponse{
				FileName:  bannermodel.Name,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}
		//将图片上传到七牛云
		if global.Config.QiNiu.Enable {
			filepath,err:=qiniu.UploadImage(byteData,filename,"gvb")
			if err!=nil {
				global.Log.Error(err)
				continue
			}
	
				reslist = append(reslist, UploadImagesResponse{
					FileName:  filepath,
					IsSuccess: true,
					Msg:       "图片上传七牛成功",
				})
			global.DB.Create(&models.BannerModel{
				Path: filepath,
				Hash: imageHash,
				Name: filename,
				Type: ctype.QiNiu,
			})
			continue
		}
		
		//将图片上传到指定目录
		err = c.SaveUploadedFile(file, filepath)
		if err != nil {
			global.Log.Error(err)
			reslist = append(reslist, UploadImagesResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}
		//上传成功消息提示
		reslist = append(reslist, UploadImagesResponse{
			FileName:  file.Filename,
			IsSuccess: true,
			Msg:       "图片上传成功",
		})

		//将图片保存到数据库
		global.DB.Create(&models.BannerModel{
			Path: filepath,
			Hash: imageHash,
			Name: filename,
			Type: ctype.QiNiu,
		})

	}
	response.OkWithData(reslist, c)
}
