package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"errors"
	"gorm.io/gorm"
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

func (*FollowStruct) RelationUpdate(userId int64, toUserId int64, actionType int32) (isExist bool) {
	f := dal.ConnQuery.Follow

	follow := &model.Follow{UserID: int32(userId), FunID: int32(toUserId), Removed: actionType}

	_, err := f.Where(f.UserID.Eq(follow.UserID), f.FunID.Eq(follow.FunID)).Update(f.Removed, follow.Removed)

	utils.CatchErr("更新错误", err)

	isExist = !errors.Is(err, gorm.ErrRecordNotFound)

	return isExist
}

func (*FollowStruct) RelationCreate(userId int64, toUserId int64) (err error) {
	f := dal.ConnQuery.Follow

	follow := &model.Follow{UserID: int32(userId), FunID: int32(toUserId)}

	err = f.Create(follow)

	return err
}
