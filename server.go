package main

import (
	"ByteDance/dal"
	"ByteDance/pkg/middleware"
	"ByteDance/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 启动redis
	err1 := dal.InitClient()
	if err1 != nil {
		//redis连接错误
		fmt.Println("Redis连接失败")
	} else {
		r.Use(middleware.RateMiddleware)
	}
	initRouter(r) // 初始化路由
	err := r.Run(":8000")
	utils.CatchErr("Run", err)
}
