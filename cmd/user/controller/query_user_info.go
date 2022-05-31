package controller

import (
	"ByteDance/cmd/user"
	"ByteDance/cmd/user/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 用户登录返回值
type loginResponse struct {
	common.Response
	user.LoginData
}

// 用户注册返回值
type regUserResponse struct {
	common.Response
	user.RegUserData
}

// 获取用户信息
type getUserInfoResponse struct {
	common.Response
	User user.GetUserInfoData `json:"user"`
}

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	regUserData, isExist := service.RegUser(username, password)
	if isExist {
		c.JSON(http.StatusOK, regUserResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.AlreadyRegisteredStatusMsg,
			},
		})
	} else {
		c.JSON(http.StatusOK, regUserResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  msg.RegisterSuccessStatusMsg,
			},
			RegUserData: user.RegUserData{
				ID:    regUserData.ID,
				Token: regUserData.Token,
			},
		})
	}
}

func LoginUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// state 1:登陆成功 0:账号或密码错误 -1:账号已被冻结
	loginData, state := service.LoginUser(username, password)
	if state == 1 {
		// 登陆成功
		c.JSON(http.StatusOK, loginResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  msg.LoginSuccessStatusMsg,
			},
			LoginData: user.LoginData{
				ID:    loginData.ID,
				Token: loginData.Token,
			},
		})
	} else if state == 0 {
		// 账号或密码错误
		c.JSON(http.StatusOK, loginResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.WrongUsernameOrPasswordMsg,
			},
		})
	} else if state == -1 {
		// 账号已被冻结
		c.JSON(http.StatusOK, loginResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.AccountBlocked,
			},
		})
	}
}

func GetUserInfo(c *gin.Context) {
	userIDStr := c.Query("user_id")
	token := c.Query("token")
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, getUserInfoResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
		return
	}

	fmt.Println(token) // token 暂时不管了，等中间件了

	// 根据userID查询用户名 获取关注数、粉丝数
	userInfoData, state := service.GetUserInfo(int32(userID))
	// state 0:userID不存在  1:查询成功  -1:查询失败
	if state == 0 {
		c.JSON(http.StatusOK, &getUserInfoResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.UserIDNotExistMsg,
			},
		})
	} else if state == -1 {
		c.JSON(http.StatusOK, &getUserInfoResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.GetUserInfoFailedMsg,
			},
		})
	} else {
		c.JSON(http.StatusOK,
			&getUserInfoResponse{
				Response: common.Response{
					StatusCode: 0,
					StatusMsg:  msg.GetUserInfoSuccessMsg,
				},
				User: *userInfoData,
			})
	}
}
