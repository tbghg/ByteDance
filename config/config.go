package config

import "time"

const MySQLDSN = "byte_dance:7efftEaAtzjEwfT4@tcp(106.15.107.229:3306)/byte_dance?charset=utf8mb4&parseTime=True&loc=Local"
const Salt = "bytedance"

/**
jwt
*/
const Issuer = "xhx"
const MySecret = "EfKxHcOK2bmAL683GHBrVsxY9w4E16mUVvLM4abIoB7E3uiQzYJD8+ro27K1uno8S/vHQ7euhKCEeUodeRulhnRGLUojsGwn1jtwp/czcYmSPGSAdF+EGQn5S4qrEGJhNEpbYFYPKGgC2nUpf4E9aRHMipMKiFVU5pd2ku+VKeKN/Ism7hJyAzkMmlxUy9X5N3IuXRgTIMDAiHD/HSlVVew63lwkQv7jeWJvTtft8Gb7adRjxZDOEMVY3/aSjzIeolbinwHKVGgDV2grRWekscawus2Rjew+q9EnnzqR0RpVcjtasPBWal1foFrNyroKuvF2iyEZK+aqHE5vWQ4yRA=="
const TokenExpirationTime = 2 * time.Hour * time.Duration(1)

/* 原本打算这样写的，后来还是换成const了
type MysqlConfig struct {
	user     string
	pwd      string
	ip       string
	port     string
	database string
}

func (m MysqlConfig) DSN() string {
	m.ip = "106.15.107.229"
	m.user = "byte_dance"
	m.port = "3306"
	m.pwd = "7efftEaAtzjEwfT4"
	m.database = "byte_dance"

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.user, m.pwd, m.ip, m.port, m.database)
}
*/
