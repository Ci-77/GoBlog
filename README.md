# GoBlog
基于go开发的博客系统
## 项目结构
```
GVB
├─ api      //逻辑处理api
│  ├─ advert_api
│  │  ├─ advert_create.go
│  │  ├─ advert_delete.go
│  │  ├─ advert_list.go
│  │  ├─ advert_update.go
│  │  ├─ enter.go
│  │  └─ tmp
│  │     └─ runner-build.exe
│  ├─ article_api
│  │  ├─ article_create.go
│  │  ├─ article_list.go
│  │  ├─ article_remove.go
│  │  ├─ artile_update.go
│  │  └─ enter.go
│  ├─ enter.go
│  ├─ images_api
│  │  ├─ enter.go
│  │  ├─ image_list.go
│  │  ├─ image_name_list.go
│  │  ├─ image_remove.go
│  │  ├─ image_update.go
│  │  └─ image_upload.go
│  ├─ menu_api
│  │  ├─ enter .go
│  │  ├─ menu_creat.go
│  │  ├─ menu_delete.go
│  │  ├─ menu_namelist.go
│  │  └─ menu_update.go
│  ├─ message_api
│  │  ├─ enter.go
│  │  ├─ message_create.go
│  │  └─ message_list.go
│  ├─ settings_api
│  │  ├─ enter.go
│  │  ├─ settings_email.go
│  │  ├─ settings_email_update.go
│  │  ├─ settings_info.go
│  │  └─ settings_info_update.go
│  ├─ tag_api
│  │  ├─ enter.go
│  │  ├─ tag_create.go
│  │  ├─ tag_list.go
│  │  ├─ tag_remove.go
│  │  └─ tag_update.go
│  └─ user_api
│     ├─ email_login.go
│     ├─ enter.go
│     ├─ qq_login.go
│     ├─ role_update.go
│     ├─ user_list.go
│     ├─ user_login.go
│     ├─ user_logout.go
│     ├─ user_remove.go
│     └─ user_update.go
├─ config   //yaml文件映射结构体
│  ├─ config.go
│  ├─ config_email.go
│  ├─ config_jwt.go
│  ├─ config_logger.go
│  ├─ config_mysql.go
│  ├─ config_qiniu.go
│  ├─ config_qq.go
│  ├─ config_redis.go
│  ├─ config_site_info.go
│  ├─ config_system.go
│  └─ coonfig_upload.go
├─ core   //文件读取，数据库，日志等初始化
│  ├─ config.go
│  ├─ gorm.go
│  ├─ logger.go
│  └─ redis.go
├─ docs  //swagger api接口文档生成文件
│  ├─ docs.go
│  ├─ swagger.json
│  └─ swagger.yaml
├─ flag      //终端控制
│  ├─ create_user.go
│  ├─ db.go
│  └─ enter.go
├─ global   //抽离公共变量
│  └─ global.go
├─ go.mod
├─ go.sum
├─ main.go
├─ middleware  //中间件
│  └─ jwt_auth.go
├─ models      //数据库模型映射
│  ├─ advert_models.go
│  ├─ article_model.go
│  ├─ banner_model.go
│  ├─ comment_model.go
│  ├─ ctype     //数据库模型添加枚举方法
│  │  ├─ array.go
│  │  ├─ image_type.go
│  │  ├─ role_type.go
│  │  └─ status_type.go
│  ├─ enter.go
│  ├─ fade_back_models.go
│  ├─ log_model.go
│  ├─ menu_banner_model.go
│  ├─ message_model.go
│  ├─ meun_model.go
│  ├─ response    //封装错误码包，api响应封装
│  │  ├─ err_code.go
│  │  └─ response.go
│  ├─ tag_model.go
│  ├─ user_collect_model.go
│  └─ user_model.go
├─ myapp
├─ plugins   //外部接口
│  ├─ qiniu
│  │  └─ enter.go
│  └─ qq_login
│     └─ enter.go
├─ README.md
├─ router   //路由
│  ├─ advert_router.go
│  ├─ article_router.go
│  ├─ images_router.go
│  ├─ menu_router.go
│  ├─ message_router.go
│  ├─ router.go
│  ├─ settings_router.go
│  ├─ tag_router.go
│  └─ user_router.go
├─ service //服务端添加方法
│  ├─ common
│  │  ├─ service_list.go
│  │  └─ tmp
│  │     └─ runner-build.exe
│  ├─ enter.go
│  ├─ redis_ser
│  │  └─ enter.go
│  └─ user_ser
│     └─ enter.go
├─ settings.yaml  //配置yaml文件
├─ templates //前端文件
├─ test      //单元测试
│  ├─ jwt_test.go
│  ├─ pwd_test.go
│  ├─ reids_test.go
│  └─ string_test.go
├─ tmp
│  ├─ runner-build.exe
│  └─ runner-build.exe~
├─ uploads  //上传的图片
│  ├─ avater
│  │  └─ default.jpg
│  └─ file
│     ├─ 2022-01-24-211132.jpg.jpg
│     ├─ 2022-01-24-211139.jpg
│     ├─ 20230214_230942.jpg
│     ├─ IMG_20220717_170745.jpg
│     ├─ IMG_20220717_170804.jpg
│     ├─ IMG_20220717_170826.jpg
│     ├─ IMG_20220717_170845.jpg
│     ├─ IMG_20220717_170901.jpg
│     ├─ IMG_20220717_170923.jpg
│     ├─ IMG_20220717_182104.jpg
│     └─ IMG_20220717_182118.jpg
└─ utils  //一些工具组件
   ├─ desens
   │  └─ enter.go
   ├─ jwt
   │  ├─ enter.go
   │  ├─ gen_token.go
   │  └─ parse_token.go
   ├─ pwd
   │  ├─ Md5.go
   │  ├─ pwd.go
   │  ├─ utils.go
   │  └─ valid.go
   └─ random
      ├─ code.go
      └─ string.go

```

## 项目技术栈：
gin,gorm,docker,jwt,mysql,redis,swagger

## 项目简介：
项目基于gin+gorm，通过读取yaml配置文件完成统一配置；实现了用户管理，系统管理，菜单管理，图片管理，广告管理，文章管理，信息管理等模块。基于logrus实现自定义log；基于gin实现封装统一响应；配置统一错误码包；通过gorm来实现数据库表迁移，避免sql注入；以上模块均实现基本的增加，批量删除，修改以及分页查询；在用户管理模块，实现QQ登录和邮箱登录的接口，然后通过jwt实现用户鉴权；由于是jwt携带token无法实现用户注销，将携带的token放到redis设置过期时间来实现用户注销；对查询的用户列表实现数据脱敏；在图片管理模块，实现图片上传到七牛云和本地两种方式，对上传的图片设置白名单，限制大小并通过检索MD5生成的hash实现上传图片的唯一性；在信息管理模块，通过用户id实现用户之间的信息创建。并返回信息列表；在文章管理模块，由于文章表比较复杂，涵盖一对一，一对多，多对多关系，使用了gorm的联表查询和预加载，实现文章列表的查询；本项目在完成以后引入swagger接口文章，方便与前端同学进行信息交换；最后编写Dockerfile文件，生成docker镜像，提交到github，与阿里云完成关联，存储到阿里云私人镜像仓库