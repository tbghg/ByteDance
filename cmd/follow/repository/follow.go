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

func (*FollowStruct) RelationUpdate(userId int32, toUserId int32, actionType int32) (RowsAffected int64) {
	f := dal.ConnQuery.Follow

	follow := &model.Follow{UserID: toUserId, FunID: userId, Removed: actionType}

	var removed int32

	if actionType == 2 {
		//取消关注 removed为1
		removed = 1
	} else {
		//关注 removed为0
		removed = 0
	}

	row, err := f.Where(f.UserID.Eq(follow.UserID), f.FunID.Eq(follow.FunID)).Update(f.Removed, removed)

	utils.CatchErr("更新错误", err)

	return row.RowsAffected
}

func (*FollowStruct) RelationCreate(userId int32, toUserId int32) (err error) {
	f := dal.ConnQuery.Follow

	follow := &model.Follow{UserID: toUserId, FunID: userId}

	err = f.Create(follow)

	return err
}

func (*FollowStruct) GetFollowById(userId int32) ([]model.Follow, error) {

	f := dal.ConnQuery.Follow
	var result []model.Follow
	err := f.Select(f.UserID).Where(f.FunID.Eq(userId), f.Removed.Eq(0), f.Deleted.Eq(0)).Scan(&result)

	utils.CatchErr("获取关注列表id错误", err)

	return result, err
}

func (*FollowStruct) GetFollowerById(userId int32) ([]model.Follow, error) {

	f := dal.ConnQuery.Follow
	var result []model.Follow
	err := f.Select(f.FunID).Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.UserID.Eq(userId)).Scan(&result)

	utils.CatchErr("获取粉丝列表id错误", err)

	return result, err
}

func (*FollowStruct) QueryIsFollowById(userId int32, funId int32) (isFollow bool) {

	f := dal.ConnQuery.Follow

	_, err := f.Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.FunID.Eq(userId), f.UserID.Eq(funId)).Take()

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询到存在相关记录
		return true
	}
	return false
}
