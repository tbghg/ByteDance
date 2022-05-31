package service

import (
	"ByteDance/cmd/comment/repository"
	"ByteDance/utils"
)

type CommentDaoRequest struct {
	UserId      int64
	Token       string
	VideoId     int64
	ActionType  int32
	CommentText string
	CommentId   int64
}

func RelationAction(userId int32, videoId int32, commentText string, actionType int32) (err error) {
	//更新 如果数据库没有该数据则返回IsExist = 0
	IsExist := repository.CommentDao.RelationUpdate(userId, videoId, actionType)

	if IsExist == 0 {
		//添加该数据
		err = repository.CommentDao.RelationCreate(userId, videoId, commentText)
		utils.CatchErr("添加失败", err)
	}

	return err
}
