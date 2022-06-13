package controller

import (
	"ByteDance/cmd/follow"
	"ByteDance/cmd/follow/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// RelationActionResponse 关注操作返回值
type RelationActionResponse struct {
	common.Response
}

// FollowListResponse 关注的所有用户列表的返回值
type FollowListResponse struct {
	common.Response
	UserList []follow.User `json:"user_list"`
}

//关注与取消请求
type FollowActionRequest struct {
	Token      string `form:"token"        validate:"required,jwt"`
	ToUserId   int64  `form:"to_user_id"   validate:"required,numeric,min=1"`
	ActionType int32  `form:"action_type"  validate:"required,numeric,oneof=1 2"`
}

//关注、粉丝列表请求
type ListRequest struct {
	UserId int64  `form:"user_id" validate:"required,numeric,min=1"`
	Token  string `form:"token"   validate:"required,jwt"`
}

/**
关注操作
*/
func RelationAction(c *gin.Context) {
	var r FollowActionRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	//获取token中的userid
	value, _ := c.Get("user_id")
	UserId, _ := value.(int)

	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
			return
		}
	}

	err = service.RelationAction(int32(UserId), int32(r.ToUserId), int32(r.ActionType))

	if err != nil {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: -1}})
		return
	}
	if r.ActionType == 1 {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.FollowSuccessMsg}})
	} else {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.UnFollowSuccessMsg}})
	}
}

// FollowList 获取登录用户关注的所有用户列表
func FollowList(c *gin.Context) {
	var r ListRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
			return
		}
	}

	UserList, err := service.GetFollowListById(r.UserId)

	if err == nil {
		c.JSON(http.StatusOK, FollowListResponse{common.Response{
			StatusCode: 0,
			StatusMsg:  msg.GetFollowUserListSuccessMsg},
			UserList})
	} else {
		c.JSON(http.StatusOK, FollowListResponse{common.Response{StatusCode: -1, StatusMsg: msg.GetFollowUserListFailedMsg}, UserList})
	}
}

// FollowerList 获取登录用户关注的粉丝用户列表
func FollowerList(c *gin.Context) {
	var r ListRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
			return
		}
	}

	UserList, err := service.GetFollowerListById(r.UserId)

	if err == nil {
		c.JSON(http.StatusOK, FollowListResponse{common.Response{
			StatusCode: 0,
			StatusMsg:  msg.GetFollowerUserListSuccessMsg},
			UserList})
	} else {
		c.JSON(http.StatusOK, FollowListResponse{common.Response{StatusCode: -1, StatusMsg: msg.GetFollowerUserListFailedMsg}, UserList})
	}
}
