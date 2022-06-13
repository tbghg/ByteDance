# 青训营抖音项目文档

## 成员分工

+ 用户模块：田冰航、王智轶、张建行
+ 视频流模块：田冰航
+ 关注模块：徐洪湘
+ 评论模块：向政昌
+ 点赞模块：向政昌

## 功能

实现了接口文档中给出的所有接口

+ 用户模块：注册、登录、获取用户信息
+ 视频流模块：发布视频、获取Feed流、查看个人已发布视频
+ 关注模块：关注操作、获取关注列表、获取粉丝列表
+ 评论模块：评论操作、获取评论列表
+ 点赞模块：点赞操作、获取点赞列表

## 项目说明

### 说明

1. 视频模块中采用阿里云OSS对象存储，数据库部署在组员的服务器中，但服务器性能较差。
2. 采用ffmpeg获取视频封面，ffmpeg.exe已同步上传，对于windows以外的电脑需要提前安装ffmpeg。
3. redis并不是启动项目所必须的，但会缺少限制频率的功能

### 使用

数据库可以本地创建，也可使用小组部署于服务器上的，但是服务器性能较差

也可根据[表设计](#表设计)中给出的建表语句在本地创建数据库，修改ByteDance/pkg/common/config.go中的MySqlDSN，连接本地数据库

启动redis（非必须），在ByteDance目录下运行`go build && ByteDance.exe`，端口开放于8000

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

## 技术相关

### 技术栈使用

+ Gin
+ Gen
+ Mysql
+ 阿里云OSS
+ Redis
+ JWT

### 表设计

共含有user、video、follow、favorite、comment五个表

![image-20220613044339573](https://s2.loli.net/2022/06/13/BarFNPSl3JHu6Io.png)

+ user表

  字段：id、username(UK)、password、enable、deleted、login_time、create_time

  索引：username,deleted联合索引

```mysql
create table user
(
    id          int auto_increment comment 'PK，直接自增'
        primary key,
    username    varchar(32)                        not null comment 'UK，账号',
    password    varchar(32)                        not null comment '密码（MD5）',
    enable      tinyint  default 1                 null comment '账号是否可用',
    deleted     tinyint  default 0                 null comment '删除标识位',
    login_time  datetime default CURRENT_TIMESTAMP null,
    create_time datetime default CURRENT_TIMESTAMP null comment '注册时间'
)
    comment '用户表，储存用户信息';

create index user_username_deleted_index
    on user (username, deleted);
```



+ video表
  字段：id、author_id(FK)、play_url、cover_url、time、title、removed、deleted
  索引：
  1. author_id,removed,deleted联合索引
  2. time,removed,deleted联合索引

```mysql
create table video
(
    id        int auto_increment
        primary key,
    author_id int               not null,
    play_url  varchar(32)       not null,
    cover_url varchar(32)       not null,
    time      int               not null,
    title     varchar(128)      not null,
    removed   tinyint default 0 not null,
    deleted   tinyint default 0 not null,
    constraint video_user_id_fk
        foreign key (author_id) references user (id)
)
    comment '存储视频信息';

create index video_author_id_removed_deleted_index
    on video (author_id, removed, deleted);

create index video_time_removed_deleted_index
    on video (time, removed, deleted);
```

+ follow表
  字段：id、user_id(FK)、fun_id(FK)、removed、deleted
  索引：
  1. user_id,removed,deleted联合索引
  2. fun_id,removed,deleted联合索引
  3. 设计目的：便于快速统计关注数、粉丝数

```mysql
create table follow
(
    id      int auto_increment
        primary key,
    user_id int               null,
    fun_id  int               not null,
    removed tinyint default 0 not null,
    deleted tinyint default 0 not null,
    constraint follow_user_id2fun_fk_2
        foreign key (fun_id) references user (id),
    constraint follow_user_id2user_fk
        foreign key (user_id) references user (id)
)
    comment '关注表';

create index follow_fun_id_removed_deleted_index
    on follow (fun_id, removed, deleted);

create index follow_user_id_removed_deleted_index
    on follow (user_id, removed, deleted);

```



+ favorite表
  字段：id、video_id(FK)、user_id(FK)、removed、deleted
  索引：
  1. video_id,removed,deleted联合索引
  2. user_id,removed,deleted联合索引

```mysql
create table favorite
(
    id       int auto_increment
        primary key,
    video_id int                not null,
    user_id  int                not null,
    removed  tinyint default -1 not null,
    deleted  tinyint default 0  not null,
    constraint favorite_user_id_fk
        foreign key (user_id) references user (id),
    constraint favorite_video_id_fk
        foreign key (video_id) references video (id)
)
    comment '用户视频点赞表';

create index favorite_user_id_video_id_removed_deleted_index
    on favorite (user_id, video_id, removed, deleted);

create index favorite_video_id_removed_deleted_user_id_index
    on favorite (video_id, removed, deleted, user_id);
```

+ comment表
  字段：id、user_id、video_id、create_time、removed、deleted、content
  索引：
  1. user_id
  2. video_id,removed,deleted联合索引
  3. create_time,removed,deleted联合索引

```mysql
create table comment
(
    id          int auto_increment
        primary key,
    user_id     int                                not null,
    video_id    int                                not null,
    create_time datetime default CURRENT_TIMESTAMP not null,
    removed     tinyint  default 0                 not null,
    deleted     tinyint  default 0                 not null,
    content     text                               not null,
    constraint comment_user_id_fk
        foreign key (user_id) references user (id),
    constraint comment_video_id_fk
        foreign key (video_id) references video (id)
)
    comment '评论表';

create index comment_create_time_removed_deleted_index
    on comment (create_time, removed, deleted);

create index comment_video_id_removed_deleted_index
    on comment (video_id, removed, deleted);
```



### 技术细节

#### 视频投稿

采用阿里云OSS对象存储，通过ffmpeg获取视频第一帧作为封面，使用雪花算法生成文件名。

##### Gen

小组人员进行表设计，部署在服务器上，通过Gen生成模型和查询方式。对于常用查询方法采用自定义方法的方式，简化代码且减少转化时间。

#### Token验证

采用JWT进行Token验证，中间件中对Token进行解析和判断，Token不合法的请求被阻止，Token合法的请求则将token中携带的user_id信息设置到 *gin.Context 声明的参数中供后续使用，考虑到前端未采用refresh_token进行刷新，将Token过期时间设置为14天。

#### Validator参数校验

采用validator对请求参数进行合法性校验，不符合参数要求的请求将被阻止，避免SQL注入，防止通过发送请求对程序进破坏。

#### Redis请求频率限制

采用Redis记录请求ip，设置过期时间为1秒，同一ip1秒内访问超过100次的请求在中间件中被阻止，避免网站负载升高或者造成网站带宽阻塞而拒绝或无法响应正常用户的请求，防止通过发送请求对程序进破坏。