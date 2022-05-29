package service

import (
	"ByteDance/cmd/user"
	"ByteDance/cmd/user/repository"
	"ByteDance/utils"
)

func RegUser(username string, password string) (regUserData *user.RegUserData, isExist bool) {
	// 检测账号是否重复，但是接口文档并没有指出账号重复如何返回
	isExist = repository.UserDao.IsUsernameExist(username)
	if !isExist {
		id, err := repository.UserDao.CreateUser(username, password)
		utils.CatchErr("CreateUser", err)
		token, err := utils.GenToken(id)
		utils.CatchErr("tokenError", err)
		regUserData = &user.RegUserData{
			ID:    id,
			Token: token,
		}
	}
	return regUserData, isExist
}

func LoginUser(username string, password string) (*user.LoginData, bool) {
	// 出于安全考虑未区分账号不存在、密码不正确的情况
	id, isCorrect := repository.UserDao.CheckPassword(username, password)
	if !isCorrect {
		return nil, isCorrect
	} else {
		token, err := utils.GenToken(id)
		utils.CatchErr("tokenError", err)
		loginData := &user.LoginData{
			ID:    id,
			Token: token,
		}
		return loginData, isCorrect
	}
}
