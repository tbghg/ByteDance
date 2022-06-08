package main

import (
	"ByteDance/pkg/middleware"
	"ByteDance/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 启动redis
	err1 := middleware.InitClient()
	if err1 != nil {
		//redis连接错误
		panic(err1)
	}
	fmt.Println("Redis连接成功")

	r := gin.Default()
	initRouter(r) // 初始化路由
	err := r.Run(":8000")
	utils.CatchErr("Run", err)
}
