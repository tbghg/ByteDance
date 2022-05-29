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

func LoginUser(username string, password string) (loginData *LoginData, isCorrect bool) {
	// 1. 传参传账号、密码，查询只查username，select 俩删除标识位、密码，判断是否删除\封禁，以及密码是否相等
	// 2. 登录失败则直接返回isCorrect
	// 3. 登陆成功就获取token，并返回
	return loginData, isCorrect // 待完善
}
