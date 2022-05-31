package middleware

import (
	"ByteDance/pkg/msg"
	"ByteDance/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var MySecret = utils.MySecret

/* JwtMiddleware jwt中间件
使用方法：路由组最后use(utils.JwtMiddleware 参考favorite路由组)
*/
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		tokenStr := c.Query("token")
		token, err := jwt.ParseWithClaims(tokenStr, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return MySecret, nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 { //token格式错误
					c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg.TokenValidationErrorMalformed})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //token过期
					c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg.TokenValidationErrorExpired})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { //token未激活
					c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg.TokenValidationErrorNotValidYet})
					c.Abort() //阻止执行
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg.TokenHandleFailed})
					c.Abort() //阻止执行
					return
				}
			}
		}

		if _, ok := token.Claims.(*utils.MyClaims); ok && token.Valid {
			c.Next()
			return
		}
		//失效的token
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": msg.TokenValid})
		c.Abort() //阻止执行
		return
	}
}
