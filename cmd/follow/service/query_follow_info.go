package service

import (
	"ByteDance/cmd/follow"
	"ByteDance/cmd/follow/repository"
	"ByteDance/dal/method"
	"ByteDance/utils"
)

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

func GetFollowListById(userId int64) (userList []follow.User, err error) {

	//根据登录用户id获取关注用户id
	followList, err := repository.FollowDao.GetFollowById(int32(userId))

	//根据FollowList的长度初始化UserList

	userList = make([]follow.User, len(followList))

	for index, followData := range followList {
		//根据关注用户id查Name

		user := method.QueryUserById(followData.FunID)
		//根据关注用户id查关注总数和粉丝总数
		followerCount, followCount, _ := method.QueryFollowCount(followData.FunID)

		userList[index] = follow.User{
			Id:            int64(followData.FunID),
			Name:          user.Username,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      true,
		}
	}
	return userList, err
}
