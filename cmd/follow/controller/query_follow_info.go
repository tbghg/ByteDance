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

type RelationRequest struct {
	UserId     int64
	Token      string
	ToUserId   int64
	ActionType int32
}

func RelationAction(c *gin.Context) {
	userIdtStr := c.Query("user_id")
	token := c.Query("token")
	println(token)
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type")

	userId, err := strconv.ParseInt(userIdtStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
	}
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
	}

	err = service.RelationAction(userId, toUserId, int32(actionType))

	if err != nil {
		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: 0}})
	} else {
		if actionType == 1 {
			c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.FollowSuccessMsg}})
		}

		c.JSON(http.StatusOK, RelationActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.UnFollowSuccessMsg}})
	}

}
