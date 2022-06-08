package middleware

import (
	"ByteDance/pkg/common"
	"ByteDance/pkg/msg"
	"ByteDance/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var mySecret = []byte(common.MySecret)

//点赞操作返回值
type MiddlewareResponse struct {
	common.Response
}

/* JwtMiddleware jwt中间件
使用方法：路由组最后use(utils.JwtMiddleware 参考favorite路由组)
*/
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
					c.JSON(http.StatusOK, MiddlewareResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.TokenValidationErrorMalformed}})

					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //token过期
					c.JSON(http.StatusOK, MiddlewareResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.TokenValidationErrorExpired}})
					c.Abort() //阻止执行
					return
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { //token未激活
					c.JSON(http.StatusOK, MiddlewareResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.TokenValidationErrorNotValidYet}})
					c.Abort() //阻止执行
					return
				} else {
					c.JSON(http.StatusOK, MiddlewareResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.TokenHandleFailed}})
					c.Abort() //阻止执行
					return
				}
			}
		}

		if claims, ok := token.Claims.(*utils.MyClaims); ok && token.Valid {
			id := claims.ID
			fmt.Println(id)
			c.Set("user_id", id)

			c.Next()
			return
		}
		//失效的token
		c.JSON(http.StatusOK, MiddlewareResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.TokenValid}})
		c.Abort() //阻止执行
		return
	}
}

var redisDb *redis.Client

// 连接到redis
func InitClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// ip限流中间件
func RateMiddleware(c *gin.Context) {
	// 10 秒刷新key为IP(c.ClientIP())的r值为0
	err := redisDb.SetNX(c.ClientIP(), 0, 10*time.Second).Err()

	// 每次访问，这个IP的对应的值加一
	redisDb.Incr(c.ClientIP())
	if err != nil {
		panic(err)
	}

	// 获取IP访问的次数
	var val int
	val, err = redisDb.Get(c.ClientIP()).Int()
	if err != nil {
		panic(err)
	}
	// 如果大于100次，返回403
	if val > 100 {
		c.Abort()
		c.JSON(http.StatusForbidden, MiddlewareResponse{Response: common.Response{StatusCode: -1, StatusMsg: msg.RequestTooFastErrorMsg}})
		return
	} else {
		// 到下一个中间件
		c.Next()
	}
}
