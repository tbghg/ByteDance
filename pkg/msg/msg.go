package msg

import "time"

// AlreadyRegisteredStatusMsg 注册
const AlreadyRegisteredStatusMsg = "该用户名已被注册"
const RegisterSuccessStatusMsg = "注册成功"

//jwt
const Issuer = "xhx"
const MySecret = "EfKxHcOK2bmAL683GHBrVsxY9w4E16mUVvLM4abIoB7E3uiQzYJD8+ro27K1uno8S/vHQ7euhKCEeUodeRulhnRGLUojsGwn1jtwp/czcYmSPGSAdF+EGQn5S4qrEGJhNEpbYFYPKGgC2nUpf4E9aRHMipMKiFVU5pd2ku+VKeKN/Ism7hJyAzkMmlxUy9X5N3IuXRgTIMDAiHD/HSlVVew63lwkQv7jeWJvTtft8Gb7adRjxZDOEMVY3/aSjzIeolbinwHKVGgDV2grRWekscawus2Rjew+q9EnnzqR0RpVcjtasPBWal1foFrNyroKuvF2iyEZK+aqHE5vWQ4yRA=="
const TokenExpirationTime = 2 * time.Hour * time.Duration(1)
const TokenValidationErrorMalformed = "token 格式错误"
const TokenValidationErrorExpired = "token 已经过期"
const TokenValidationErrorNotValidYet = "token 尚未激活"
const TokenValid = "无效 token"
const TokenHandleFailed = "无法处理此Token"
