package controller

import (
	"ByteDance/cmd/comment"
	"ByteDance/cmd/comment/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//评论操作返回
type CommentActionResponse struct {
	common.Response
}

//评列表返回
type CommentListResponse struct {
	common.Response
	CommentList []comment.TheCommentInfo `json:"comment_list"`
}

//评论与取消请求
type CommentActionRequest struct {
	Token       string `form:"token"         validate:"required,jwt"`
	VideoId     int64  `form:"video_id"      validate:"required,numeric,min=1"`
	ActionType  int32  `form:"action_type"   validate:"required,numeric,oneof=1 2"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

//评论列表请求
type CommentListRequest struct {
	VideoId int64  `form:"video_id" validate:"required,numeric,min=1"`
	Token   string `form:"token"    validate:"required,jwt"`
}

//评论及取消评论操作
func CommentAction(c *gin.Context) {

	var r CommentActionRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)

	//数据检测
	//评论
	if r.ActionType == 1 && len(r.CommentText) == 0 {
		c.JSON(http.StatusBadRequest, CommentActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.CommentFailedMsg}})
		return
	}
	//取消评论
	if r.ActionType == 2 && r.CommentId <= 0 {
		c.JSON(http.StatusBadRequest, CommentActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.CommentFailedMsg}})
		return
	}

	//获取token中的userid
	value, _ := c.Get("user_id")
	UserId, _ := value.(int)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, CommentActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
			return
		}
	}
	err = service.CommentAction(int32(UserId), int32(r.VideoId), r.CommentText, int32(r.CommentId))

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.CommentFailedMsg}})
		return
	}

	if r.ActionType == 1 {
		c.JSON(http.StatusOK, CommentActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.CommentSuccessMsg}})
	} else {
		c.JSON(http.StatusOK, CommentActionResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.UnCommentSuccessMsg}})
	}

}

//评论列表
func CommentList(c *gin.Context) {
	var r CommentListRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, CommentActionResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
			return
		}
	}
	commentInfo, _ := service.CommentList(int32(r.VideoId))
	//获取成功
	c.JSON(http.StatusOK, &CommentListResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  msg.GetCommentUserListSuccessMsg,
		},
		CommentList: commentInfo,
	})

}
