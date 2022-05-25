package main

import (
	"ByteDance/cmd/user/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	GRoute := r.Group("/douyin")
	{
		user := GRoute.Group("/user")
		{
			user.POST("/register/", controller.RegisterUser)
		}
	}
}