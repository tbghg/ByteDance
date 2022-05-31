package main

import (
	favoriteController "ByteDance/cmd/favorite/controller"
	relationController "ByteDance/cmd/follow/controller"
	userController "ByteDance/cmd/user/controller"
	"ByteDance/utils"

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
			user.POST("/login/", userController.LoginUser)
			user.GET("/", userController.GetUserInfo)
		}
		//follow路由组
		relation := GRoute.Group("relation")
		{
			relation.POST("/action/", relationController.RelationAction)
		}
		//favorite路由组
		favorite := GRoute.Group("/favorite").Use(utils.JwtMiddleware())
		{

			favorite.POST("/action/", favoriteController.FavoriteAction)
			favorite.GET("/list/", favoriteController.FavoriteList)
		}

	}
}
