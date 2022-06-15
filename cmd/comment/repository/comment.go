package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"sync"
)

// CommentInfo 接收comment表返回数据
type CommentInfo struct {
	ID            int32
	UserID        int32
	Username      string
	Content       string
	CreateDate    string
	FavoriteCount int
	IsFavorite    bool
}

type CommentStruct struct {
}

var (
	CommentDao  *CommentStruct
	commentOnce sync.Once
)

// 单例模式
func init() {
	commentOnce.Do(func() {
		CommentDao = &CommentStruct{}
	})
}

// CommentUpdate 取消评论
func (*CommentStruct) CommentUpdate(commentId int32) (RowsAffected int64) {
	c := dal.ConnQuery.Comment

	comment := &model.Comment{ID: commentId}

	row, err := c.Where(c.ID.Eq(comment.ID)).Update(c.Removed, 1)

	if err != nil {
		utils.Log.Error("评论更新错误" + err.Error())
	}

	return row.RowsAffected
}

// CommentCreate 评论操作
func (*CommentStruct) CommentCreate(userId int32, videoId int32, commentText string) bool {
	c := dal.ConnQuery.Comment

	comment := &model.Comment{UserID: userId, VideoID: videoId, Content: commentText}

	err := c.Create(comment)
	if err != nil {
		utils.Log.Error("插入评论表错误" + err.Error())
		return false
	}
	return true
}

// CommentList 评论列表
func (*CommentStruct) CommentList(videoId int32) ([]CommentInfo, bool) {
	c := dal.ConnQuery.Comment
	u := dal.ConnQuery.User

	var result []CommentInfo
	// 内联查询
	err := c.Select(c.ID, c.UserID, u.Username, c.Content, c.CreateTime.As("CreateDate")).Where(c.VideoID.Eq(videoId), c.Removed.Eq(0), c.Deleted.Eq(0)).Join(u, u.ID.EqCol(c.UserID)).Order(c.ID.Desc()).Scan(&result)
	if err != nil {
		utils.Log.Error("获取评论错误" + err.Error())
		return nil, false
	}
	return result, true
}
