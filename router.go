package main

import (
	relationController "ByteDance/cmd/follow/controller"
	userController "ByteDance/cmd/user/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// GRoute总路由组
	GRoute := r.Group("/douyin")
	{
		// user路由组
		user := GRoute.Group("/user")
		{
			user.POST("/register/", userController.RegisterUser)
		}
		//follow路由组
		relation := GRoute.Group("relation")
		{
			relation.POST("/action/", relationController.RelationAction)
		}

	}
}
