package controller

import (
	"ByteDance/cmd/favorite/repository"
	"ByteDance/cmd/favorite/service"
	"ByteDance/cmd/video"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// FavoriteListResponse 点赞列表返回值
type FavoriteListResponse struct {
	common.Response
	VideoList []video.TheVideoInfo `json:"video_list"`
}

// FavoriteActionRequest 点赞与取消请求
type FavoriteActionRequest struct {
	Token      string `form:"token" validate:"required"`
	VideoId    int64  `form:"video_id" validate:"required,numeric"`
	ActionType int32  `form:"action_type" validate:"required,numeric"`
}

// FavoriteListRequest 点赞列表请求
type FavoriteListRequest struct {
	UserId int64  `form:"user_id" validate:"required,numeric"`
	Token  string `form:"token" validate:"required"`
}

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	var r FavoriteActionRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	//获取token中的userid
	value, success := c.Get("user_id")
	if !success {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.TokenParameterAcquisitionError,
			})
		return
	}
	userID, _ := value.(int)

	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg})
			return
		}
	}

	success = repository.FavoriteDao.FavoriteAction(int32(userID), int32(r.VideoId))

	if !success {
		c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.FavoriteFailedMsg})
		return
	}
	if r.ActionType == 1 {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: msg.FavoriteSuccessMsg})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: msg.UnFavoriteSuccessMsg})
	}

}

// FavoriteList 点赞列表
func FavoriteList(c *gin.Context) {
	var r FavoriteListRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg})
			return
		}
	}
	videoInfo, _ := service.FavoriteList(int32(r.UserId))

	c.JSON(http.StatusOK, &FavoriteListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  msg.GetFavoriteUserListSuccessMsg,
		},
		VideoList: videoInfo,
	})
}
