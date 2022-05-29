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

func RelationAction(userId int64, toUserId int64, actionType int32) (err error) {
	//更新 如果数据库没有该数据则返回IsExist = false
	IsExist := repository.FollowDao.RelationUpdate(userId, toUserId, actionType)

	if !IsExist {
		//添加该数据
		err = repository.FollowDao.RelationCreate(userId, toUserId)
		utils.CatchErr("添加失败", err)
	}

	return err
}
