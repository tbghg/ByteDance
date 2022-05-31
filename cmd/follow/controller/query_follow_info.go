package controller

import (
	"ByteDance/cmd/follow/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RelationActionResponse struct {
	common.Response
}

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
