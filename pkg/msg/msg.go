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
	UserIDNotExistMsg = "用户ID不存在"
	GetFailedMsg      = "获取用户信息失败"
	GetSuccessMsg     = "获取用户信息成功"
)

// JWT
const (
	TokenValidationErrorMalformed   = "token 格式错误"
	TokenValidationErrorExpired     = "token 已经过期"
	TokenValidationErrorNotValidYet = "token 尚未激活"
	TokenValid                      = "无效 token"
	TokenHandleFailed               = "无法处理此Token"
)

//Follow
const (
	FollowSuccessMsg   = "关注成功"
	UnFollowSuccessMsg = "取消关注成功"
)

//favorite
const (
	FavoriteSuccessMsg   = "点赞成功"
	UnFavoriteSuccessMsg = "取消点赞成功"
)

//common
const (
	DataFormatErrorMsg = "数据格式错误"
)
