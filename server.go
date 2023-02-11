package main

import (
	"ByteDance/dal"
	"ByteDance/pkg/middleware"
	"ByteDance/utils"
	"github.com/gin-gonic/gin"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	configInit(wg)

	r := gin.Default()

	// 启动redis
	if dal.InitClient() {
		r.Use(middleware.RateMiddleware)
	}
	initRouter(r) // 初始化路由
	err := r.Run(":8000")
	if err != nil {
		wg.Wait()
		utils.Log.Fatal("服务启动失败 " + err.Error())
	}
}

func configInit(wg *sync.WaitGroup) {
	utils.LogConfig() // 初始化日志配置
	go func() {
		utils.OSSInit() // OSS初始化，将ConnQuery与数据库绑定
		defer wg.Done()
	}()
	go func() {
		dal.MySQLInit() // MySQL初始化，连接数据库
		defer wg.Done()
	}()
	go func() {
		utils.SensitiveWordInit() // 读取敏感词，写入内存中
		defer wg.Done()
	}()
}
