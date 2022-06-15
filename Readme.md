# 青训营抖音项目文档

## 项目说明

### 实现功能

实现了接口文档中给出的所有接口

+ 用户模块：注册、登录、获取用户信息
+ 视频流模块：发布视频、获取`Feed`流、查看个人已发布视频
+ 关注模块：关注操作、获取关注列表、获取粉丝列表
+ 评论模块：评论操作、获取评论列表
+ 点赞模块：点赞操作、获取点赞列表

### 环境配置

1. Go版本>=1.17.3
2. 数据库：MySQL8.0
3. Redis：3.2.100

### 项目使用

1. 已将数据库部署于服务器上，也可根据[表设计](https://yvrcskowz5.feishu.cn/docs/doccnJpAemQe5YEr9TmIxL2JCXb#表设计)模块中给出的建表语句在本地创建数据库
2. 启动`Redis`（非必须）
3. 在`ByteDance/pkg/common/config.go`中填写相应配置项（也可使用当前默认配置）
4. 安装依赖。在`ByteDance`目录下运行`go mod tidy`
5. 运行。运行`go build && ByteDance.exe`，端口开放于`8000`

### 项目说明

1. 视频模块中采用阿里云`OSS`对象存储
2. 数据库部署在服务器中，但服务器性能较差
3. 采用`ffmpeg`获取视频封面，`ffmpeg.exe`已同步上传项目，但对于`windows`以外的电脑需要提前安装`ffmpeg`
4. `Redis`并不是启动项目所必须的，但缺省时会缺少限制频率的功能

### 项目结构

```
ByteDance
│  .gitignore
│  ffmpeg.exe	// 截取视频第一帧
│  go.mod
│  go.sum
│  Readme.md
│  router.go	// 创建路由
│  server.go	// 项目启动入口
│
├─cmd
│  ├─user
│  │  │  user_common_model.go	// user模块中共用的结构体
│  │  │
│  │  ├─controller	// 控制层，接受参数，编写流程逻辑，返回信息
│  │  │      query_user_info.go
│  │  │
│  │  ├─repository	// 负责与数据库的交互
│  │  │      user.go
│  │  │
│  │  └─service		// 处理流程中的主要函数
│  │          query_user_info.go
│  ├─comment	// 其他模块与user模块结构相同
│  ├─favorite
│  ├─follow
│  └─video
│
├─dal
│  │  dal.go	// 初始化，将ConnQuery与数据库绑定
│  │
│  ├─method		// 自定义查询方法，用Gen生成
│  │      method.go
│  │
│  ├─model		// Gen生成的数据模型
│  │      comment.gen.go
│  │      favorite.gen.go
│  │      follow.gen.go
│  │      user.gen.go
│  │      video.gen.go
│  │
│  └─query		// Gen生成的数据库操作方法
│          comment.gen.go
│          favorite.gen.go
│          follow.gen.go
│          gen.go
│          user.gen.go
│          video.gen.go
│
├─logs		// 日志存放
├─pkg		
│  ├─common
│  │      common.go		// 模块公用部分
│  │      config.go		// 配置项
│  │
│  ├─middleware			// 中间件
│  │      middleware.go
│  │
│  └─msg
│          msg.go		// 定义返回消息
│
└─utils		// 工具类
    │  catchErr.go		// 捕捉错误
    │  jwt.go			// 生成Token令牌
    │  md5.go			// md5加密
    │  snowflake.go		// 雪花算法
    │  upload_file.go	// OSS中上传文件
    │
    └─generate
            generate.go	// Gen生成模块与方法
```



## 成员分工

| **成员** |                           **分工**                           |
| :------: | :----------------------------------------------------------: |
|  田冰航  | 数据库设计，项目结构设计，用户注册功能，获取视频流功能，上传视频功能，查看已发布视频功能 |
|  向政昌  | Validate数据验证，实现redis中间件限制频率，评论功能， 点赞功能你，相关功能文档撰写 |
|  徐洪湘  | JWT令牌功能实现，数据库设计，项目结构设计，关注功能，相关功能文档攥写 |
|  王智轶  |                获取用户信息，相关功能文档撰写                |
|  张建行  |                用户登录功能，相关功能文档撰写                |

## 后记

技术相关及功能实现请移步汇报文档：[极简版抖音项目汇报文档（打工魂小组）](https://yvrcskowz5.feishu.cn/docs/doccnJpAemQe5YEr9TmIxL2JCXb#)