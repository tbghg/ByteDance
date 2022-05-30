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

func (*UserDaoStruct) CheckPassword(username string, password string) (id int, state int) {
	u := dal.ConnQuery.User
	user, err := u.Where(u.Username.Eq(username)).Take()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 查询到存在相关记录
		if user.Deleted == 0 && utils.Md5(password) == user.Password {
			// 密码正确
			if user.Enable == 1 {
				// 账号可以使用
				state = 1
				id = int(user.ID)
			} else {
				// 账号无法使用
				state = -1
			}
		}
	} else {
		// 不存在相关记录，检查err是否为空，不为空说明出现错误
		utils.CatchErr("登录查询错误", err)
	}
	// 账号密码错误时id、state默认初始值为0
	return id, state
}
