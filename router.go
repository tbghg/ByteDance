package main

import (
	"ByteDance/cmd/user/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// GRoute总路由组
	GRoute := r.Group("/douyin")
	{
		// user路由组
		user := GRoute.Group("/user")
		{
			user.POST("/register/", controller.RegisterUser)
		}

	}
}
