package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"sync"
)

type Favorite struct {
	model.Comment
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

func (*CommentStruct) RelationUpdate(userId int32, videoId int32, actionType int32) (RowsAffected int64) {
	f := dal.ConnQuery.Comment

	comment := &model.Comment{UserID: userId, VideoID: videoId, Removed: actionType}

	row, err := f.Where(f.UserID.Eq(comment.UserID), f.VideoID.Eq(comment.VideoID)).Update(f.Removed, comment.Removed)

	utils.CatchErr("更新错误", err)

	return row.RowsAffected
}

func (*CommentStruct) RelationCreate(userId int32, videoId int32, commentText string) (err error) {
	f := dal.ConnQuery.Comment

	comment := &model.Comment{UserID: userId, VideoID: videoId, Content: commentText}

	err = f.Create(comment)

	return err
}
