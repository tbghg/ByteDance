package main

import (
	"ByteDance/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initRouter(r) // 初始化路由
	err := r.Run(":8000")
	utils.CatchErr("Run", err)
}
