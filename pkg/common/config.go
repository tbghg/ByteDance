package common

import "time"

const MySqlDSN = "byte_dance:7efftEaAtzjEwfT4@tcp(106.15.107.229:3306)/byte_dance?charset=utf8mb4&parseTime=True&loc=Local"

// MD5Salt MD5加密时的盐
const MD5Salt = "UII34HJ6OIO"

// JWT
const (
	Issuer              = "xhx" // 签发人
	MySecret            = "Fy3Jfa5AD"
	TokenExpirationTime = 14 * 24 * time.Hour * time.Duration(1) // Token过期时间
)

// OSSPreURL OSS前缀
const OSSPreURL = "https://byte-dance-01.oss-accelerate.aliyuncs.com/video/"
