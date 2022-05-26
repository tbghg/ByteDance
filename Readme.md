未完待整理

```
ByteDance
│  .gitignore	// 根据自己的需要往里加
│  go.mod
│  go.sum
│  Readme.md
│  router.go	// 创建路由，所有的路由将会在这里创建
│  server.go	// 项目启动入口
│
├─cmd
│  │
│  ├─user  // 用户模块
│  │  ├─controller  // 控制层，接受参数，编写流程逻辑（但是流程中的主要功能在service中实现），返回信息
│  │  │      query_user_info.go
│  │  │
│  │  ├─repository  // 负责与数据库的交互（与数据库相关的所有交互均在这里进行）
│  │  │      user.go
│  │  │
│  │  └─service     // 处理流程中的主要函数，从repository中调用数据库操作（一般service中不出现数据库相关操作）
│  │          query_user_info.go
│  │
│  │  // 以下模块均仿照user的风格进行书写
│  ├─follow         // 关注\粉丝
│  ├─comment        // 评论
│  ├─favorite       // 赞
│  └─video          // 视频
│
│
├─dal
│  │  dal.go
│  │
│  ├─method     // 可在此处自定义查询方法然后用gen生成，如果打算使用请务必在群里说一声（详细使用方法参考gen）
│  │      method.go
│  │
│  ├─model      // gen生成的数据模型
│  │      comment.gen.go
│  │      favorite.gen.go
│  │      follow.gen.go
│  │      user.gen.go
│  │      video.gen.go
│  │
│  └─query      // gen生成的数据库操作方法
│          comment.gen.go
│          favorite.gen.go
│          follow.gen.go
│          gen.go
│          user.gen.go
│          video.gen.go
│
├─logs      // 日志存放处
│
├─pkg
│  ├─common // 模块共用的结构体或函数，以及配置项
│  │      common.go
│  │
│  ├─middleware     // 中间件
│  └─msg    // 定义返回的message消息
│          msg.go
│
└─utils     // 工具类
    │  catchErr.go      // 捕获错误的方法
    │  jwt.go           // 生成\解析JWT
    │  md5.go           // 进行MD5加密
    │
    └─generate  // 根据数据库生成模型及操作
            generate.go
```