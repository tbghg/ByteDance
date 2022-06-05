package controller

import (
	"ByteDance/cmd/favorite/service"
	"ByteDance/cmd/video"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//点赞操作返回值
type FavoriteActionResponse struct {
	common.Response
}

//点赞列表返回值
type FavoriteListResponse struct {
	common.Response
	VideoList []video.TheVideoInfo `json:"video_list"`
}

//点赞与取消请求
type FavoriteActionRequest struct {
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
	//获取token中的userid
	value, _ := c.Get("user_id")
	UserId, _ := value.(int)

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

	err = service.RelationAction(int32(UserId), int32(r.VideoId))

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
	var r FavoriteListRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回，测试用，上线删除
			c.JSON(http.StatusBadRequest, gin.H{
				"翻译前": errors.Error(),
				"翻译后": errors.Translate(common.Trans),
			})
			return
		}
	}
	videoInfo, _ := service.RelationList(int32(r.UserId))
	//获取成功
	c.JSON(http.StatusOK, &FavoriteListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  msg.GetVideoInfoSuccessMsg,
		},
		VideoList: videoInfo,
	})
}
