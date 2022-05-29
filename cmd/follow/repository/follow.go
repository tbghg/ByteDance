package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"sync"
)

type Follow struct {
	model.Follow
}

type FollowStruct struct {
}

var (
	FollowDao  *FollowStruct
	followOnce sync.Once
)

// 单例模式
func init() {
	followOnce.Do(func() {
		FollowDao = &FollowStruct{}
	})
}

func (*FollowStruct) RelationUpdate(userId int32, toUserId int32, actionType int32) (RowsAffected int64) {
	f := dal.ConnQuery.Follow

	follow := &model.Follow{UserID: userId, FunID: toUserId, Removed: actionType}

	row, err := f.Where(f.UserID.Eq(follow.UserID), f.FunID.Eq(follow.FunID)).Update(f.Removed, follow.Removed)

	utils.CatchErr("更新错误", err)

	return row.RowsAffected
}

func (*FollowStruct) RelationCreate(userId int32, toUserId int32) (err error) {
	f := dal.ConnQuery.Follow

	follow := &model.Follow{UserID: userId, FunID: toUserId}

	err = f.Create(follow)

	return err
}
