package controller

import (
	"ByteDance/cmd/favorite/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type FavoriteActionResponse struct {
	common.Response
}

//点赞与取消请求
type FavoriteActionRequest struct {
	UserId     int64  `form:"user_id" validate:"required,numeric"`
	Token      string `form:"token" validate:"required"`
	VideoId    int64  `form:"video_id" validate:"required,numeric"`
	ActionType int32  `form:"action_type" validate:"required,numeric"`
}

//点赞列表请求
type FavoriteListRequest struct {
	UserId int64  `form:"user_id" validate:"required,numeric"`
	Token  string `form:"token" validate:"required"`
}

//点赞操作
func FavoriteAction(c *gin.Context) {
	var r FavoriteActionRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)

	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, gin.H{
				"翻译前": errors.Error(),
				"翻译后": errors.Translate(common.Trans),
			})
			return
		}
	}
	//userIdtStr := c.Query("user_id")
	//videoIdStr := c.Query("video_id")
	//actionTypeStr := c.Query("action_type")

	//userId, err := strconv.ParseInt(userIdtStr, 10, 64)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
	//	return
	//}
	//videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
	//	return
	//}
	//actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	//
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.DataFormatErrorMsg}})
	//	return
	//}

	err = service.RelationAction(int32(r.UserId), int32(r.VideoId), int32(r.ActionType))

	if err != nil {
		c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: -1}})
		return
	}
	if r.ActionType == 1 {
		c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.FavoriteSuccessMsg}})
	} else {
		c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.UnFavoriteSuccessMsg}})
	}

}

//点赞列表
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, FavoriteActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: "放行成功！"}})
}
