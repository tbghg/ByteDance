package dal

import (
	"ByteDance/dal/query"
	"ByteDance/pkg/common"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var ConnQuery *query.Query
var once sync.Once

// 初始化，将ConnQuery与数据库绑定
func init() {
	once.Do(func() {
		ConnQuery = getQueryConnection()
	})
}

func getQueryConnection() *query.Query {
	var err error
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(common.MySqlDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	ConnQuery = query.Use(db)
	return ConnQuery
}
