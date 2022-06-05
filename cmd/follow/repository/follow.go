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

	follow := &model.Follow{UserID: userId, FunID: toUserId}

	err = f.Create(follow)

	return err
}

func (*FollowStruct) GetFollowById(userId int32) (followList []*model.Follow, err error) {

	f := dal.ConnQuery.Follow

	followList, err = f.Select(f.FunID).Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.UserID.Eq(userId)).Find()

	utils.CatchErr("获取关注列表id错误", err)

	return followList, err
}
