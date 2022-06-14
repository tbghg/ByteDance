package middleware

import (
	"ByteDance/dal"
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"ByteDance/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var mySecret = []byte(common.MySecret)

// JwtMiddleware jwt中间件 使用方法：路由组最后use(utils.JwtMiddleware 参考favorite路由组)
func JwtMiddleware(method string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取token
		var tokenStr string
		if method == "query" {
			tokenStr = c.Query("token")
		} else if method == "form-data" {
			tokenStr = c.PostForm("token")
		} else if method == "feed" {
			tokenStr = c.Query("token")
		}

		token, err := jwt.ParseWithClaims(tokenStr, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		})
		if err != nil {
			if method == "feed" {
				c.Next()
				return
			}
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 { //token格式错误
					c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.TokenValidationErrorMalformed})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //token过期
					c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.TokenValidationErrorExpired})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { //token未激活
					c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.TokenValidationErrorNotValidYet})
					c.Abort() //阻止执行
					return
				} else {
					c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.TokenHandleFailed})
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
		c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.TokenValid})
		c.Abort() //阻止执行
		return
	}
}

// RateMiddleware ip限流中间件
// ip限流中间件
func RateMiddleware(c *gin.Context) {
	// 5 秒刷新key为IP(c.ClientIP())的r值为0
	err := dal.RedisDb.SetNX(c.ClientIP(), 0, 5*time.Second).Err()

	// 每次访问，这个IP的对应的值加一
	dal.RedisDb.Incr(c.ClientIP())
	if err != nil {
		panic(err)
	}

	// 获取IP访问的次数
	var val int
	val, err = dal.RedisDb.Get(c.ClientIP()).Int()
	if err != nil {
		panic(err)
	}
	// 如果5秒内大于100次
	if val > 100 {
		c.Abort()
		c.JSON(http.StatusOK, common.Response{StatusCode: -1, StatusMsg: msg.RequestTooFastErrorMsg})
		return
	} else {
		// 到下一个中间件
		c.Next()
	}
}
