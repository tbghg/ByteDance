package main

import (
	"ByteDance/dal/method"
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

	g.ApplyBasic(g.GenerateModel("user"))
	g.ApplyBasic(g.GenerateModel("video"))
	g.ApplyInterface(func(method method.FavoriteMethod) {}, g.GenerateModel("favorite"))
	g.ApplyInterface(func(method method.CommentMethod) {}, g.GenerateModel("comment"))
	g.ApplyInterface(func(method method.FollowMethod) {}, g.GenerateModel("follow"))
	g.Execute()
}
