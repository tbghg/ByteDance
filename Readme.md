未完待整理

```
ByteDance
│  .gitignore	// 根据自己的需要往里加
│  go.mod
│  go.sum
│  Readme.md
│  router.go	// 创建路由
│  server.go	// 项目启动入口
│
├─cmd
│  └─user       //user服务的业务代码
│      ├─controller
│      │      register_user.go
│      │
│      ├─repository
│      │      user.go
│      │
│      └─service
│              register_user.go
│
├─config	// 配置项
│      config.go
│
├─dal
│  │  dal.go
│  │
│  ├─method	// 如果需要自定义方法的话可以用，不过需要重新生成了
│  │      method.go
│  │
│  ├─model	// gen框架生成，模型，可在repository中调用
│  │      comment.gen.go
│  │      favorite.gen.go
│  │      follow.gen.go
│  │      user.gen.go
│  │      video.gen.go
│  │
│  └─query	// gen框架生成，CURD，可在repository中调用
│          comment.gen.go
│          favorite.gen.go
│          follow.gen.go
│          gen.go
│          user.gen.go
│          video.gen.go
│
├─logs		// 日志
├─pkg
│  ├─common	
│  │      common.go
│  │
│  ├─errno
│  ├─middleware	// 中间件
│  └─msg	// 定义返回的message消息
│          msg.go
│
└─utils		// 工具类层
    │  catchErr.go
    │  jwt.go
    │  md5.go
    │
    └─generate	// gen，生成模型和CURD方法
            generate.go
```

