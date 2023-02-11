use byte_dance;

create table user
(
    id          int auto_increment comment 'PK，直接自增'
        primary key,
    username    varchar(32)                         not null comment 'UK，账号',
    password    varchar(32)                         not null comment '密码（MD5）',
    enable      tinyint   default 1                 null comment '账号是否可用',
    deleted     tinyint   default 0                 null comment '删除标识位',
    login_time  datetime  default CURRENT_TIMESTAMP null,
    create_time timestamp default CURRENT_TIMESTAMP null comment '注册时间'
)
    comment '用户表，储存用户信息';

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

create index user_username_deleted_index
    on user (username, deleted);

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

create index video_author_id_removed_deleted_index
    on video (author_id, removed, deleted);

create index video_time_removed_deleted_index
    on video (time, removed, deleted);

