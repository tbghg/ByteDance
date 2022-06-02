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

//关注操作返回值
type RelationActionResponse struct {
	common.Response
}

//关注的所有用户列表的返回值
type FollowListResponse struct {
	common.Response
	UserList []follow.User `json:"user_list"`
}

/**
关注操作
*/
func RelationAction(c *gin.Context) {
	userIdtStr := c.Query("user_id")
	token := c.Query("token")
	println("token:" + token)
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type")

	userId, err := strconv.ParseInt(userIdtStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
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

/**
获取登录用户关注的所有用户列表
*/
func FollowList(c *gin.Context) {
	userIdtStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdtStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}

	UserList, err := service.GetFollowListById(userId)

	if err == nil {
		c.JSON(http.StatusOK, FollowListResponse{common.Response{
			StatusCode: 0,
			StatusMsg:  msg.GetFollowUserListSuccessMsg},
			UserList})
	} else {
		c.JSON(http.StatusBadRequest, FollowListResponse{common.Response{StatusCode: -1, StatusMsg: msg.GetFollowUserListFailedMsg}, UserList})
	}
}
