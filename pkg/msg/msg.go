package msg

// 用户注册
const (
	AlreadyRegisteredStatusMsg = "该用户名已被注册"
	RegisterSuccessStatusMsg   = "注册成功"
	MatchFailedStatusMsg       = "账号密码需要小于32字符，密码包含至少一位数字，字母和特殊字符"
)

// 用户登录
const (
	WrongUsernameOrPasswordMsg = "用户名或密码错误"
	LoginSuccessStatusMsg      = "登陆成功"
	AccountBlocked             = "账号已被冻结"
)

// 获取用户信息
const (
	UserIDNotExistMsg     = "用户ID不存在"
	GetUserInfoFailedMsg  = "获取用户信息失败"
	GetUserInfoSuccessMsg = "获取用户信息成功"
)

// video
const (
	HasNoVideoMsg            = "已经没有视频了"
	GetVideoInfoSuccessMsg   = "获取视频信息成功"
	PublishVideoFailedMsg    = "上传视频失败"
	PublishVideoSuccessMsg   = "上传视频成功"
	GetPublishListFailedMsg  = "获取已发布视频失败"
	GetPublishListSuccessMsg = "获取已发布视频成功"
)

// JWT
const (
	TokenValidationErrorMalformed   = "token 格式错误"
	TokenValidationErrorExpired     = "登录状态已失效，请重新登录"
	TokenValidationErrorNotValidYet = "token 尚未激活"
	TokenValid                      = "无效 token"
	TokenHandleFailed               = "无法处理此Token"
	TokenParameterAcquisitionError  = "token参数获取错误"
)

//Follow
const (
	FollowSuccessMsg              = "关注成功"
	UnFollowSuccessMsg            = "取消关注成功"
	GetFollowUserListSuccessMsg   = "获取用户关注列表成功"
	GetFollowUserListFailedMsg    = "获取用户关注列表失败"
	GetFollowerUserListSuccessMsg = "获取粉丝列表成功"
	GetFollowerUserListFailedMsg  = "获取粉丝列表失败"
)

//Favorite
const (
	FavoriteSuccessMsg            = "点赞成功"
	UnFavoriteSuccessMsg          = "取消点赞成功"
	GetFavoriteUserListSuccessMsg = "获取用户点赞列表成功"
	FavoriteFailedMsg             = "操作失败"
)

//Comment
const (
	CommentSuccessMsg            = "评论成功"
	UnCommentSuccessMsg          = "取消评论成功"
	GetCommentUserListSuccessMsg = "获取视频评论列表成功"
	CommentFailedMsg             = "操作失败"
)

//common
const (
	DataFormatErrorMsg = "请求数据缺失或格式错误"
)

//IP限速
const (
	RequestTooFastErrorMsg = "请求过快，稍后再试"
)
