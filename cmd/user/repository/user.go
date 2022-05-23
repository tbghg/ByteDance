package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"errors"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	model.User
}

type UserDaoStruct struct {
}

var (
	UserDao  *UserDaoStruct
	userOnce sync.Once
)

// 单例模式
func init() {
	userOnce.Do(func() {
		UserDao = &UserDaoStruct{}
	})
}

func (*UserDaoStruct) IsUsernameExist(username string) (isExist bool) {
	u := dal.ConnQuery.User
	_, err := u.Where(u.Username.Eq(username)).Take()
	isExist = !errors.Is(err, gorm.ErrRecordNotFound)
	return isExist
}

func (*UserDaoStruct) CreateUser(username string, password string) (int, error) {
	u := dal.ConnQuery.User
	user := &model.User{Username: username, Password: utils.Md5(password)}

	err := u.Create(user)

	return int(user.ID), err
}
