package controller

import (
	"ByteDance/cmd/user"
	"ByteDance/cmd/user/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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

//注册 登录请求
type RegisterLoginRequest struct {
	Username string `form:"username"  validate:"required"`
	Password string `form:"password"  validate:"required"`
}

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var r RegisterLoginRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, regUserResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
			return
		}
	}
	//密码强度检测
	match := common.MatchStr(r.Password)
	if !match || len(r.Username) > 32 {
		c.JSON(http.StatusOK, regUserResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.MatchFailedStatusMsg,
			},
		})
		return
	}

	regUserData, isExist := service.RegUser(r.Username, r.Password)
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
	var r RegisterLoginRequest
	// 接收参数并绑定
	err := c.ShouldBindQuery(&r)
	// 使用common包中Validate验证器
	err = common.Validate.Struct(r)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			// 翻译，并返回
			c.JSON(http.StatusBadRequest, loginResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.DataFormatErrorMsg}})
			return
		}
	}
	// state 1:登陆成功 0:账号或密码错误 -1:账号已被冻结
	loginData, state := service.LoginUser(r.Username, r.Password)
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

	userID, success := c.Get("user_id")
	if !success {
		c.JSON(http.StatusOK,
			common.Response{
				StatusCode: -1,
				StatusMsg:  msg.TokenParameterAcquisitionError,
			})
		return
	}
	// 根据userID查询用户名 获取关注数、粉丝数
	userInfoData, success2 := service.GetUserInfo(int32(userID.(int)))
	// state 0:userID不存在  1:查询成功  -1:查询失败
	if !success2 {
		c.JSON(http.StatusOK, &getUserInfoResponse{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.UserIDNotExistMsg,
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
