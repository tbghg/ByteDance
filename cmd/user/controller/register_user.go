package controller

import (
	"ByteDance/cmd/user/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"ByteDance/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegData struct {
	common.Response
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

func RegisterUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	regUserData, isExist := service.RegUser(username, password)

	if isExist {
		c.JSON(http.StatusOK, RegData{
			Response: common.Response{
				StatusCode: -1,
				StatusMsg:  msg.AlreadyRegisteredStatusMsg,
			},
		})
	}
	// jwt根据id生成token，略
	token, err := utils.GenToken(username)
	utils.CatchErr("tokenError", err)
	c.JSON(http.StatusOK, RegData{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  msg.RegisterSuccessStatusMsg,
		},
		UserID: regUserData.ID,
		Token:  token, // jwt生成，中间数据感觉只传id就行，暂时不想写了，要挂科了
	})
}
