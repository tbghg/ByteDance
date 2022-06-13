package controller

import (
	"ByteDance/cmd/follow"
	"ByteDance/cmd/follow/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

// RelationAction 关注操作
func RelationAction(c *gin.Context) {
	userIdStr := c.Query("user_id")
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	actionType, err2 := strconv.ParseInt(actionTypeStr, 10, 64)

	if err != nil || err2 != nil {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}

	err = service.RelationAction(int32(userId), int32(toUserId), int32(actionType))

	if err != nil {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: -1}})
		return
	}
	if actionType == 1 {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.FollowSuccessMsg}})
	} else {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.UnFollowSuccessMsg}})
	}
}

// FollowList 获取登录用户关注的所有用户列表
func FollowList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	UserList, err := service.GetFollowListById(int64(userId))

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
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	UserList, err := service.GetFollowerListById(int64(userId))

	if err == nil {
		c.JSON(http.StatusOK, FollowListResponse{common.Response{
			StatusCode: 0,
			StatusMsg:  msg.GetFollowerUserListSuccessMsg},
			UserList})
	} else {
		c.JSON(http.StatusOK, FollowListResponse{common.Response{StatusCode: -1, StatusMsg: msg.GetFollowerUserListFailedMsg}, UserList})
	}
}
