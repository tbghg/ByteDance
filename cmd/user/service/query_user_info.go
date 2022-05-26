package service

import (
	"ByteDance/cmd/user/repository"
	"ByteDance/utils"
)

type RegUserData struct {
	ID    int    `json:"user_id"`
	Token string `json:"token"`
}

func RegUser(username string, password string) (regUserData *RegUserData, isExist bool) {
	// 检测账号是否重复，但是接口文档并没有指出账号重复如何返回
	isExist = repository.UserDao.IsUsernameExist(username)
	if !isExist {
		id, err := repository.UserDao.CreateUser(username, password)
		utils.CatchErr("CreateUser", err)
		// jwt根据id生成token，略
		token, err := utils.GenToken(id)
		utils.CatchErr("tokenError", err)
		regUserData = &RegUserData{
			ID:    id,
			Token: token,
		}
	}
	return regUserData, isExist
}
