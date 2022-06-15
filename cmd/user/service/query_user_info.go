package service

import (
	"ByteDance/cmd/user"
	"ByteDance/cmd/user/repository"
	"ByteDance/utils"
)

func RegUser(username string, password string) (regUserData *user.RegUserData, isExist bool) {
	// 检测账号是否重复
	isExist = repository.UserDao.IsUsernameExist(username)
	if !isExist {
		id := repository.UserDao.CreateUser(username, password)
		token := utils.GenToken(id)
		regUserData = &user.RegUserData{
			ID:    id,
			Token: token,
		}
	}
	return regUserData, isExist
}

func LoginUser(username string, password string) (loginData *user.LoginData, state int) {
	// 出于安全考虑未区分账号不存在、密码不正确的情况
	var id int
	id, state = repository.UserDao.CheckPassword(username, password)
	if state == 1 {
		// 登录成功，创建loginData、计算token
		token := utils.GenToken(id)
		loginData = &user.LoginData{
			ID:    id,
			Token: token,
		}
	}
	return loginData, state
}

func GetUserInfo(userID int32) (userInfoData *user.GetUserInfoData, success bool) {

	username, followCount, followerCount, isExist := repository.UserDao.QueryUserInfoByID(userID)
	if !isExist {
		// userID 不存在相关记录
		return nil, false
	} else {
		userInfoData = &user.GetUserInfoData{
			ID:            userID,
			UseName:       username,
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      false,
		}
		return userInfoData, true
	}
}
