package main

import (
	"ByteDance/pkg/common"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db, _ = gorm.Open(mysql.Open(common.MySqlDSN))

// generate code
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../dal/query",
		ModelPkgPath: "../../dal/model",
		Mode:         gen.WithoutContext,
	})
	g.UseDB(db)

	// 如果想要自定义方法就在ByteDance/dal/model/method.go中添加相应的接口及实现，如下
	// g.ApplyInterface(func(method model.UserMethod) {},g.GenerateModel("user"))

	g.ApplyBasic(g.GenerateModel("user"))
	g.ApplyBasic(g.GenerateModel("video"))
	g.ApplyBasic(g.GenerateModel("favorite"))
	g.ApplyBasic(g.GenerateModel("comment"))
	g.ApplyBasic(g.GenerateModel("follow"))
	g.Execute()
}
