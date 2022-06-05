package msg

// 用户注册
const (
	AlreadyRegisteredStatusMsg = "该用户名已被注册"
	RegisterSuccessStatusMsg   = "注册成功"
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
	HasNoVideoMsg          = "已经没有视频了"
	GetVideoInfoSuccessMsg = "获取视频信息成功"
	PublishVideoFailedMsg  = "上传视频失败"
	PublishVideoSuccessMsg = "上传视频成功"
)

// JWT
const (
	TokenValidationErrorMalformed   = "token 格式错误"
	TokenValidationErrorExpired     = "token 已经过期"
	TokenValidationErrorNotValidYet = "token 尚未激活"
	TokenValid                      = "无效 token"
	TokenHandleFailed               = "无法处理此Token"
	TokenParameterAcquisitionError  = "token参数获取错误"
)

//Follow
const (
	FollowSuccessMsg            = "关注成功"
	UnFollowSuccessMsg          = "取消关注成功"
	GetFollowUserListSuccessMsg = "获取用户关注列表成功"
	GetFollowUserListFailedMsg  = "获取用户关注列表失败"
)

//Favorite
const (
	FavoriteSuccessMsg   = "点赞成功"
	UnFavoriteSuccessMsg = "取消点赞成功"
)

//Comment
const (
	CommentSuccessMsg   = "评论成功"
	UnCommentSuccessMsg = "取消评论成功"
)

//common
const (
	DataFormatErrorMsg = "数据格式错误"
)
