package main

import (
	favoriteController "ByteDance/cmd/favorite/controller"
	relationController "ByteDance/cmd/follow/controller"
	userController "ByteDance/cmd/user/controller"
	videoController "ByteDance/cmd/video/controller"

	commentController "ByteDance/cmd/comment/controller"
	"ByteDance/pkg/common"
	"ByteDance/pkg/middleware"
	"github.com/gin-gonic/gin"
	zhs "github.com/go-playground/validator/v10/translations/zh"
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
		favorite := GRoute.Group("/favorite").Use(middleware.JwtMiddleware())
		{
			favorite.POST("/action/", favoriteController.FavoriteAction)
			favorite.GET("/list/", favoriteController.FavoriteList)
		}
		// 视频流接口
		GRoute.GET("/feed/", videoController.GetVideoFeed)
		//comment路由组
		comment := GRoute.Group("/comment").Use(middleware.JwtMiddleware())
		{

			comment.POST("/action/", commentController.CommentAction)
			comment.GET("/list/", commentController.CommentList)
		}

	}
	// 注册翻译器
	_ = zhs.RegisterDefaultTranslations(common.Validate, common.Trans)
}
