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

func LoginUser(username string, password string) (loginData *user.LoginData, state int) {
	// 出于安全考虑未区分账号不存在、密码不正确的情况
	var id int
	id, state = repository.UserDao.CheckPassword(username, password)
	if state == 1 {
		// 登录成功，创建loginData、计算token
		token, err := utils.GenToken(id)
		utils.CatchErr("tokenError", err)
		loginData = &user.LoginData{
			ID:    id,
			Token: token,
		}
	}
	return loginData, state
}
