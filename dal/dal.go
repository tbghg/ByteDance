package dal

import (
	"ByteDance/dal/query"
	"ByteDance/pkg/common"
	"ByteDance/utils"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var ConnQuery *query.Query
var once sync.Once

// MySQLInit 初始化，将ConnQuery与数据库绑定
func MySQLInit() {
	once.Do(func() {
		ConnQuery = getQueryConnection()
	})
}

func getQueryConnection() *query.Query {
	var err error
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(common.MySqlDSN))
	if err != nil {
		utils.Log.Fatal("数据库连接错误" + err.Error())
	} else {
		utils.Log.Info("MySQL连接成功")
	}
	ConnQuery = query.Use(db)
	return ConnQuery
}

// RedisDb redis
var RedisDb *redis.Client

// InitClient 连接到redis
func InitClient() bool {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     common.RedisLocalhost, // redis地址
		Password: common.RedisPassword,  // redis密码，没有则留空
		DB:       common.RedisDB,        // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err := RedisDb.Ping().Result()
	if err != nil {
		utils.Log.Warn("Redis连接失败")
		return false
	} else {
		utils.Log.Info("Redis连接成功")
		return true
	}
}
