package controller

import (
	"ByteDance/cmd/user/service"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
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

	c.JSON(http.StatusOK, RegData{
		Response: common.Response{
			StatusCode: 0,
			StatusMsg:  msg.RegisterSuccessStatusMsg,
		},
		UserID: regUserData.ID,
		Token:  regUserData.Token,
	})
}
