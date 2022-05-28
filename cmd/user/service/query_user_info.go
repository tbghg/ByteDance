package service

import (
	"ByteDance/cmd/user/repository"
	"ByteDance/utils"
)

type RegUserData struct {
	ID    int    `json:"user_id"`
	Token string `json:"token"`
}

type LoginData struct {
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

func LoginUser(username string, password string) (*LoginData, bool) {
	// 出于安全考虑未区分账号不存在、密码不正确的情况
	id, isCorrect := repository.UserDao.CheckPassword(username, password)
	if !isCorrect {
		return nil, isCorrect
	} else {
		token, err := utils.GenToken(id)
		utils.CatchErr("tokenError", err)
		loginData := &LoginData{
			ID:    id,
			Token: token,
		}
		return loginData, isCorrect
	}
}
