package service

import (
	"ByteDance/cmd/follow"
	"ByteDance/cmd/follow/repository"
	"ByteDance/dal/method"
	"ByteDance/utils"
	"sync"
)

var (
	username      string
	followerCount int64
	followCount   int64
	isFollow      bool
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

	//根据登录用户id获取关注用户id列表

	followList, err := repository.FollowDao.GetFollowById(int32(userId))

	//根据FollowList的长度初始化UserList

	userList = make([]follow.User, len(followList))

	for index, followData := range followList {

		//可以使用并发执行
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			//根据关注用户id查用户名
			user := method.QueryUserById(followData.UserID)
			username = user.Username
		}()

		go func() {
			defer wg.Done()
			//根据关注用户id查关注总数和粉丝总数
			followCount, followerCount, _ = method.QueryFollowCount(followData.FunID)
		}()

		wg.Wait()

		userList[index] = follow.User{
			Id:            int64(followData.FunID),
			Name:          username,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      true,
		}
	}
	return userList, err
}

func GetFollowerListById(userId int64) (userList []follow.User, err error) {
	//根据登录用户id获取粉丝id列表

	followerList, err := repository.FollowDao.GetFollowerById(int32(userId))

	//根据FollowList的长度初始化UserList

	userList = make([]follow.User, len(followerList))

	for index, followData := range followerList {

		//可以使用并发执行
		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			//根据粉丝id查用户名
			user := method.QueryUserById(followData.FunID)
			username = user.Username
		}()

		go func() {
			defer wg.Done()
			//根据粉丝id查关注总数和粉丝总数
			followCount, followerCount, _ = method.QueryFollowCount(followData.UserID)
		}()

		go func() {
			defer wg.Done()
			//根据登录用户id查询是否关注该粉丝用户(互关)
			isFollow = repository.FollowDao.QueryIsFollowById(int32(userId), followData.UserID)

		}()

		wg.Wait()

		userList[index] = follow.User{
			Id:            int64(followData.UserID),
			Name:          username,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      isFollow,
		}
	}
	return userList, err
}
