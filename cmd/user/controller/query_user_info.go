package controller

import (
	"ByteDance/cmd/user/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户登录返回值
type loginResponse struct {
	common.Response
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

// 用户注册返回值
type regUserResponse struct {
	common.Response
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
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
	}

	c.JSON(http.StatusOK, regUserResponse{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  msg.RegisterSuccessStatusMsg,
		},
		UserID: regUserData.ID,
		Token:  regUserData.Token,
	})
}

func LoginUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	loginData, isCorrect := service.LoginUser(username, password)
	if !isCorrect {
		// 账号或密码错误
		c.JSON(http.StatusOK, loginResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.WrongUsernameOrPasswordMsg,
			},
		})
	} else {
		// 登陆成功
		c.JSON(http.StatusOK, loginResponse{
			Response: common.Response{
				StatusCode: 0,
				StatusMsg:  msg.LoginSuccessStatusMsg,
			},
			UserID: loginData.ID,
			Token:  loginData.Token,
		})
	}
}
