package controller

import (
	"ByteDance/cmd/favorite/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavoriteActionResponse struct {
	common.Response
}

type FavoriteRequest struct {
	UserId     int64
	Token      string
	VideoId    int64
	ActionType int32
}

//点赞操作
func FavoriteAction(c *gin.Context) {
	userIdtStr := c.Query("user_id")
	token := c.Query("token")
	println("token:" + token)
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")

	userId, err := strconv.ParseInt(userIdtStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}

	err = service.RelationAction(int32(userId), int32(videoId), int32(actionType))

	if err != nil {
		c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: -1}})
		return
	}
	if actionType == 1 {
		c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.FavoriteSuccessMsg}})
	} else {
		c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.UnFavoriteSuccessMsg}})
	}

}

//点赞列表
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: "放行成功！"}})
}
