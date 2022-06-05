package controller

import (
	"ByteDance/cmd/comment/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type CommentListResponse struct {
	common.Response
}

//评论与取消请求
type CommentActionRequest struct {
	UserId      int64  `form:"user_id" validate:"required,numeric"`
	Token       string `form:"token" validate:"required"`
	VideoId     int64  `form:"video_id" validate:"required,numeric"`
	ActionType  int32  `form:"action_type" validate:"required,numeric"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

//评论列表请求
type CommentListRequest struct {
	VideoId int64  `form:"video_id" validate:"required,numeric"`
	Token   string `form:"token" validate:"required"`
}

//评论操作
func CommentAction(c *gin.Context) {

	var r CommentActionRequest
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
	err = service.RelationAction(int32(r.UserId), int32(r.VideoId), r.CommentText, r.ActionType)

	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{Response: common.Response{StatusCode: -1}})
		return
	}

	if r.ActionType == 1 {
		c.JSON(http.StatusOK, CommentListResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.CommentSuccessMsg}})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{Response: common.Response{StatusCode: 0, StatusMsg: msg.UnCommentSuccessMsg}})
	}

}

//评论列表
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{Response: common.Response{StatusCode: 0, StatusMsg: "获取评论列表成功！"}})
}
