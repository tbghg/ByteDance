package service

import (
	"ByteDance/cmd/follow/repository"
	"ByteDance/utils"
)

type RelationRequestS struct {
	UserId     int64
	Token      string
	ToUserId   int64
	ActionType int32
}

func RelationAction(userId int32, toUserId int32, actionType int32) (err error) {
	//更新 如果数据库没有该数据则返回IsExist = 0
	IsExist := repository.FollowDao.RelationUpdate(userId, toUserId, actionType)

	if IsExist == 0 {
		//添加该数据
		err = repository.FollowDao.RelationCreate(userId, toUserId)
		utils.CatchErr("添加失败", err)
	}

	return err
}
