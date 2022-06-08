package middleware

import (
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"ByteDance/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var mySecret = []byte(common.MySecret)

// JwtMiddleware jwt中间件 使用方法：路由组最后use(utils.JwtMiddleware 参考favorite路由组)
func JwtMiddleware(method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		var tokenStr string
		if method == "query" {
			tokenStr = c.Query("token")
		} else {
			tokenStr = c.PostForm("token")
		}

		token, err := jwt.ParseWithClaims(tokenStr, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 { //token格式错误
					c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": msg.TokenValidationErrorMalformed})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //token过期
					c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": msg.TokenValidationErrorExpired})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { //token未激活
					c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": msg.TokenValidationErrorNotValidYet})
					c.Abort() //阻止执行
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": msg.TokenHandleFailed})
					c.Abort() //阻止执行
					return
				}
			}
		}

		if claims, ok := token.Claims.(*utils.MyClaims); ok && token.Valid {
			id := claims.ID
			c.Set("user_id", id)
			c.Next()
			return
		}
		//失效的token
		c.JSON(http.StatusOK, gin.H{"status_code": -1, "status_msg": msg.TokenValid})
		c.Abort() //阻止执行
		return
	}
}
