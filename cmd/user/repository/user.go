package repository

import (
	"ByteDance/dal"
	"ByteDance/dal/model"
	"ByteDance/utils"
	"errors"
	"gorm.io/gorm"
	"sync"
	"time"
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
				_, updateErr := u.Where(u.ID.Eq(user.ID)).Update(u.LoginTime, time.Now().Format("2006-01-02 15:04:05"))
				utils.CatchErr("登录时间更新", updateErr)
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

func (*UserDaoStruct) QueryUsernameByID(userID int32) (username string, isExist bool) {
	u := dal.ConnQuery.User
	user, err := u.Select(u.Username).Where(u.ID.Eq(userID)).Take()
	isExist = !errors.Is(err, gorm.ErrRecordNotFound)
	if isExist == false {
		return "", isExist
	} else {
		return user.Username, isExist
	}
}

func (*UserDaoStruct) QueryFollowCount(userID int32) (followerCount int64, followCount int64, success bool) {
	// follower_count 粉丝数	follow_count 关注数
	var err error
	f := dal.ConnQuery.Follow

	followerCount, err = f.Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.UserID.Eq(userID)).Count()
	utils.CatchErr("查询粉丝数错误", err)
	if err != nil {
		return 0, 0, false
	}

	followCount, err = f.Where(f.Deleted.Eq(0), f.Removed.Eq(0), f.FunID.Eq(userID)).Count()
	utils.CatchErr("查询关注数错误", err)
	if err != nil {
		return 0, 0, false
	}

	return followerCount, followCount, true
}
