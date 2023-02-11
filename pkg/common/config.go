package common

import "time"

const MySqlDSN = "root:123456@tcp(mysql_a:3306)/byte_dance?charset=utf8mb4"

// Redis 配置
const (
	RedisLocalhost = "redis_a:6379"
	RedisPassword  = ""
	RedisDB        = 0
)

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

// SensitiveWordsPath 敏感词路径
const SensitiveWordsPath = "./utils/SensitiveWords.txt"
