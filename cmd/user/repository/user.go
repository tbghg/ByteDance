package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"errors"
	"gorm.io/gorm"
	"sync"
)

// User 数据库查询完毕后将查询结构放在此处
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

func (*UserDaoStruct) CheckPassword(username string, password string) (int, bool) {
	u := dal.ConnQuery.User
	user, err := u.Where(u.Username.Eq(username)).Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 未查询到相关法记录，即不存在该账号
		return int(user.ID), false
	}
	if user.Deleted == 0 && user.Enable == 1 && utils.Md5(password) == user.Password {
		// 密码正确，且账号可以正常使用
		return int(user.ID), true
	} else {
		// 密码不正确，或账号被封禁
		return int(user.ID), false
	}
}
