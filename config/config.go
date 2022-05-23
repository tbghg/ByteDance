package config

const MySQLDSN = "byte_dance:7efftEaAtzjEwfT4@tcp(106.15.107.229:3306)/byte_dance?charset=utf8mb4&parseTime=True&loc=Local"
const Salt = "bytedance"

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
